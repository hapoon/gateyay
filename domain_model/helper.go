package domain_model

import (
	"testing"
	"time"
)

func ToDate(t *testing.T, date string) time.Time {
	t.Helper()
	d, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		t.Fatalf("ToDate: %v", err)
	}
	return d
}
