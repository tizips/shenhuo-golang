package helper

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

func ReadExcelRows(file *excelize.File) ([][]string, error) {
	name := file.GetSheetName(0)
	if name == "" {
		return nil, fmt.Errorf("空工作表")
	}

	rows, err := file.GetRows(name)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func HeaderIndex(headers []string) map[string]int {
	index := make(map[string]int, len(headers))
	for i, cell := range headers {
		key := strings.TrimSpace(cell)
		if key == "" {
			continue
		}
		if _, ok := index[key]; !ok {
			index[key] = i
		}
	}
	return index
}

func Cell(row []string, idx int) string {
	if idx < 0 || idx >= len(row) {
		return ""
	}
	return strings.TrimSpace(row[idx])
}

func CellByName(row []string, headers map[string]int, names ...string) string {
	for _, name := range names {
		if idx, ok := headers[name]; ok {
			return Cell(row, idx)
		}
	}
	return ""
}
