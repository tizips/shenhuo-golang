package helper

import (
	"testing"
	"time"
)

func TestPersonImportPassword(t *testing.T) {
	at := time.Date(2026, 9, 12, 10, 0, 0, 0, time.UTC)

	if got := PersonImportPassword("13812345678", at); got != "45678@20260912" {
		t.Fatalf("got %s", got)
	}

	if got := PersonImportPassword("1234", at); got != "1234@20260912" {
		t.Fatalf("got %s", got)
	}
}
