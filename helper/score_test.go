package helper

import "testing"

func TestParseScoreHeader(t *testing.T) {
	headers := []string{"姓名", "参赛号", "单位", "理论考试", "管理能力考试", "井下紧急避险考核", "自救器操作考核", "心肺复苏考核", "总分"}
	numberCol, itemCols, err := ParseScoreHeader(headers)
	if err != nil {
		t.Fatal(err)
	}
	if numberCol != 1 {
		t.Fatalf("numberCol %d", numberCol)
	}
	want := []int{3, 4, 5, 6, 7}
	for i, col := range itemCols {
		if col != want[i] {
			t.Fatalf("itemCols %v", itemCols)
		}
	}

	// 缺少固定小项列时需报错
	if _, _, err = ParseScoreHeader([]string{"参赛号", "理论考试", "总分"}); err == nil {
		t.Fatal("need all fixed items")
	}

	// 缺少参赛号列时需报错
	if _, _, err = ParseScoreHeader([]string{"姓名", "总分"}); err == nil {
		t.Fatal("need 参赛号")
	}
}

func TestScoreItemsFromRow(t *testing.T) {
	_, itemCols, err := ParseScoreHeader([]string{"参赛号", "理论考试", "管理能力考试", "井下紧急避险考核", "自救器操作考核", "心肺复苏考核"})
	if err != nil {
		t.Fatal(err)
	}

	items := ScoreItemsFromRow([]string{"001", "90", "85.5", "", "80", "95"}, itemCols)
	if len(items) != 5 {
		t.Fatalf("len %d", len(items))
	}
	if items[0].Name != "理论考试" || items[0].Value != "90" {
		t.Fatalf("%v", items[0])
	}
	if items[2].Value != "" {
		t.Fatalf("%v", items[2])
	}
	if items[4].Value != "95" {
		t.Fatalf("%v", items[4])
	}
}

func TestCalculateScoreTotal(t *testing.T) {
	total, err := CalculateScoreTotal([]string{"90", "85.5", "", "80", "95.5"})
	if err != nil {
		t.Fatal(err)
	}
	if total != "351" {
		t.Fatalf("total %s", total)
	}

	if _, err = CalculateScoreTotal([]string{"90", "abc"}); err == nil {
		t.Fatal("need error for non-numeric score")
	}
}
