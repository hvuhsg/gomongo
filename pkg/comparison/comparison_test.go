package comparison_test

import (
	"testing"

	"github.com/hvuhsg/gomongo/pkg/comparison"
)

func TestGraterAny(t *testing.T) {
	expectsTrue := []map[any]any{
		{5: 4.0},
		{3.1: 3},
		{97.1: 'a'},
		{98: 'a'},
		{0.5: 0b0},
		{"zzz": "aaa"},
		{0b1111: rune(0b1110)},
	}

	expectsFalse := []map[any]any{
		{4: 5.0},
		{2.9: 3},
		{96.9: 'a'},
		{96: 'a'},
		{0.5: 0b1111},
		{"aaa": "aaa"},
		{5.0: 5},
		{"aaa": "zzz"},
		{0b1111: rune(0b1111)},
	}

	for _, item := range expectsTrue {
		for a, b := range item {
			t.Run("expects true", func(t *testing.T) {
				isGrater, err := comparison.GraterAny(a, b)
				if err != nil || !isGrater {
					t.Errorf("Excpects %v to be grater then %v", a, b)
				}
			})
		}
	}

	for _, item := range expectsFalse {
		for a, b := range item {
			t.Run("expects false", func(t *testing.T) {
				isGrater, err := comparison.GraterAny(a, b)
				if err != nil || isGrater {
					t.Errorf("Excpects %v to be less or equal to %v", a, b)
				}
			})
		}
	}
}
