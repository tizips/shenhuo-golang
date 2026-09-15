package helper

import (
	"errors"
	"math/rand/v2"

	"github.com/tizips/shenhuo/model"
	"gorm.io/gorm"
)

func PickIDs(ids []string, n int) []string {
	if n <= 0 || len(ids) == 0 {
		return []string{}
	}

	cloned := append([]string(nil), ids...)
	rand.Shuffle(len(cloned), func(i, j int) {
		cloned[i], cloned[j] = cloned[j], cloned[i]
	})

	if n > len(cloned) {
		n = len(cloned)
	}

	return cloned[:n]
}

func DrawCategory(tx *gorm.DB, categoryID uint) ([]model.ShDraw, error) {
	var category model.ShDrawCategory
	if err := tx.First(&category, "`id`=?", categoryID).Error; err != nil {
		return nil, err
	}

	if category.Quota == 0 {
		return nil, errors.New("中签人数必须大于 0")
	}

	var drawn int64
	if err := tx.Model(&model.ShDraw{}).Where("`category_id`=?", category.ID).Count(&drawn).Error; err != nil {
		return nil, err
	}

	remain := int(category.Quota) - int(drawn)
	if remain <= 0 {
		return nil, errors.New("该类别已抽满")
	}

	var drawnIDs []string
	if err := tx.Model(&model.ShDraw{}).Where("`category_id`=?", category.ID).Pluck("person_id", &drawnIDs).Error; err != nil {
		return nil, err
	}

	query := tx.Model(&model.ShPerson{}).Select("id")
	if len(drawnIDs) > 0 {
		query = query.Where("`id` NOT IN ?", drawnIDs)
	}

	var personIDs []string
	if err := query.Pluck("id", &personIDs).Error; err != nil {
		return nil, err
	}

	picked := PickIDs(personIDs, remain)
	if len(picked) == 0 {
		return nil, errors.New("没有可抽签的人员")
	}

	draws := make([]model.ShDraw, len(picked))
	for index, id := range picked {
		draws[index] = model.ShDraw{
			CategoryID: category.ID,
			PersonID:   id,
		}
	}

	if err := tx.Create(&draws).Error; err != nil {
		return nil, err
	}

	if err := tx.Preload("Person").Preload("Category").Where("`id` IN ?", drawIDs(draws)).Find(&draws).Error; err != nil {
		return nil, err
	}

	return draws, nil
}

func drawIDs(draws []model.ShDraw) []uint {
	ids := make([]uint, len(draws))
	for index, item := range draws {
		ids[index] = item.ID
	}
	return ids
}
