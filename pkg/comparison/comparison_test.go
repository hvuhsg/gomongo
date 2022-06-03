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

func TestLesserAny(t *testing.T) {
	expectsFalse := []map[any]any{
		{5: 4.0},
		{3.1: 3},
		{97.1: 'a'},
		{98: 'a'},
		{0.5: 0b0},
		{"zzz": "aaa"},
		{0b1111: rune(0b1110)},
		{"aaa": "aaa"},
		{5.0: 5},
		{0b1111: rune(0b1111)},
	}

	expectsTrue := []map[any]any{
		{4: 5.0},
		{2.9: 3},
		{96.9: 'a'},
		{96: 'a'},
		{0.5: 0b1111},
		{"aaa": "zzz"},
	}

	for _, item := range expectsFalse {
		for a, b := range item {
			t.Run("expects false", func(t *testing.T) {
				isLesser, err := comparison.LesserAny(a, b)
				if err != nil || isLesser {
					t.Errorf("Excpects %v to be grater from %v", a, b)
				}
			})
		}
	}

	for _, item := range expectsTrue {
		for a, b := range item {
			t.Run("expects true", func(t *testing.T) {
				isLesser, err := comparison.LesserAny(a, b)
				if err != nil || !isLesser {
					t.Errorf("Excpects %v to be lesser then %v", a, b)
				}
			})
		}
	}
}

func TestEqualAny(t *testing.T) {
	expectsTrue := []map[any]any{
		{5: 5.0},
		{3.1: 3.1},
		{97: 'a'},
		{0.0: 0b0},
		{"aaa": "aaa"},
		{0b1111: rune(0b1111)},
	}

	expectsFalse := []map[any]any{
		{4: 5.0},
		{2.9: 3},
		{96.9: 'a'},
		{96: 'a'},
		{0.5: 0b1111},
		{"aab": "aaa"},
		{5.2: 5},
		{"aaz": "zzz"},
		{0b1110: rune(0b1111)},
	}

	for _, item := range expectsTrue {
		for a, b := range item {
			t.Run("expects true", func(t *testing.T) {
				isEqaule := comparison.EqualAny(a, b)
				if !isEqaule {
					t.Errorf("Excpects %v to be equal to %v", a, b)
				}
			})
		}
	}

	for _, item := range expectsFalse {
		for a, b := range item {
			t.Run("expects false", func(t *testing.T) {
				isEqaule := comparison.EqualAny(a, b)
				if isEqaule {
					t.Errorf("Excpects %v to not be equal to %v", a, b)
				}
			})
		}
	}
}
