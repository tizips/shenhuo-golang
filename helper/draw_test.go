package helper

import "testing"

func TestPickIDs(t *testing.T) {
	ids := []string{"a", "b", "c", "d"}

	got := PickIDs(ids, 2)
	if len(got) != 2 {
		t.Fatalf("len=%d", len(got))
	}

	if len(PickIDs(ids, 0)) != 0 {
		t.Fatal("n=0")
	}

	all := PickIDs(ids, 10)
	if len(all) != 4 {
		t.Fatalf("cap to source, got %d", len(all))
	}

	if len(PickIDs(nil, 3)) != 0 {
		t.Fatal("empty source")
	}
}
