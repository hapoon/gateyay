package domain_model

import (
	"testing"
)

func TestReservation(t *testing.T) {
	want := reservation{
		at:       ToDate(t, "2025-01-02 03:04:05"),
		email:    "aa@bb",
		name:     "abc",
		quantity: 1,
	}
	test := NewReservation(
		ToDate(t, "2025-01-02 03:04:05"),
		"aa@bb",
		"abc",
		1,
	)
	got := test.Equals(want)
	if !got {
		t.Error("fck")
	}
}
