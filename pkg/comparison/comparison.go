package comparison

import (
	"fmt"

	constraints "golang.org/x/exp/constraints"
)

func lesser[T constraints.Ordered](a, b T) bool {
	return a < b
}

func grater[T constraints.Ordered](a, b T) bool {
	return a > b
}

func equal(a, b any) bool {
	return a == b
}

func GraterAny(a, b any) (bool, error) {
	switch at := a.(type) {
	case int:
		switch bt := b.(type) {
		case int:
			return grater(at, bt), nil
		case float32:
			return grater(float32(at), bt), nil
		case float64:
			return grater(float64(at), bt), nil
		case rune:
			return grater(at, int(bt)), nil
		default:
			return false, fmt.Errorf("can only compere int with one of [int, float, rune]")
		}
	case float32:
		switch bt := b.(type) {
		case int:
			return grater(at, float32(bt)), nil
		case float32:
			return grater(float32(at), bt), nil
		case float64:
			return grater(float64(at), bt), nil
		case rune:
			return grater(at, float32(bt)), nil
		default:
			return false, fmt.Errorf("can only compere float with one of [int, float, rune]")
		}
	case float64:
		switch bt := b.(type) {
		case int:
			return grater(at, float64(bt)), nil
		case float32:
			return grater(float64(at), float64(bt)), nil
		case float64:
			return grater(float64(at), bt), nil
		case rune:
			return grater(at, float64(bt)), nil
		default:
			return false, fmt.Errorf("can only compere float with one of [int, float, rune]")
		}
	case rune:
		switch bt := b.(type) {
		case int:
			return grater(at, rune(bt)), nil
		case float32:
			return grater(float32(at), bt), nil
		case float64:
			return grater(float64(at), bt), nil
		case rune:
			return grater(at, bt), nil
		default:
			return false, fmt.Errorf("can only compere rune with one of [int, float, rune]")
		}
	case string:
		switch bt := b.(type) {
		case string:
			return grater(at, bt), nil
		case rune:
			return grater(at, string(bt)), nil
		default:
			return false, fmt.Errorf("can only compere string with one of [string, rune]")
		}
	case byte:
		switch bt := b.(type) {
		case byte:
			return grater(at, bt), nil
		case int:
			return grater(int(at), bt), nil
		case float32:
			return grater(float32(at), bt), nil
		case float64:
			return grater(float64(at), bt), nil
		case rune:
			return grater(float64(at), float64(bt)), nil
		default:
			return false, fmt.Errorf("can only compere byte with one of [byte, float, int, rune]")
		}
	default:
		return false, fmt.Errorf("cant convert document value to one of [int, float, complex, rune, byte]")
	}
}

func LesserAny(a, b any) (bool, error) {
	switch at := a.(type) {
	case int:
		switch bt := b.(type) {
		case int:
			return lesser(at, bt), nil
		case float32:
			return lesser(float32(at), bt), nil
		case float64:
			return lesser(float64(at), bt), nil
		case rune:
			return lesser(at, int(bt)), nil
		default:
			return false, fmt.Errorf("can only compere int with one of [int, float, rune]")
		}
	case float32:
		switch bt := b.(type) {
		case int:
			return lesser(at, float32(bt)), nil
		case float32:
			return lesser(float32(at), bt), nil
		case float64:
			return lesser(float64(at), bt), nil
		case rune:
			return lesser(at, float32(bt)), nil
		default:
			return false, fmt.Errorf("can only compere float with one of [int, float, rune]")
		}
	case float64:
		switch bt := b.(type) {
		case int:
			return lesser(at, float64(bt)), nil
		case float32:
			return lesser(float64(at), float64(bt)), nil
		case float64:
			return lesser(float64(at), bt), nil
		case rune:
			return lesser(at, float64(bt)), nil
		default:
			return false, fmt.Errorf("can only compere float with one of [int, float, rune]")
		}
	case rune:
		switch bt := b.(type) {
		case int:
			return lesser(at, rune(bt)), nil
		case float32:
			return lesser(float32(at), bt), nil
		case float64:
			return lesser(float64(at), bt), nil
		case rune:
			return lesser(at, bt), nil
		default:
			return false, fmt.Errorf("can only compere rune with one of [int, float, rune]")
		}
	case string:
		switch bt := b.(type) {
		case string:
			return lesser(at, bt), nil
		case rune:
			return lesser(at, string(bt)), nil
		default:
			return false, fmt.Errorf("can only compere string with one of [string, rune]")
		}
	case byte:
		switch bt := b.(type) {
		case byte:
			return lesser(at, bt), nil
		case int:
			return lesser(int(at), bt), nil
		case float32:
			return lesser(float32(at), bt), nil
		case float64:
			return lesser(float64(at), bt), nil
		case rune:
			return lesser(float64(at), float64(bt)), nil
		default:
			return false, fmt.Errorf("can only compere byte with one of [byte, float, int, rune]")
		}
	default:
		return false, fmt.Errorf("cant convert document value to one of [int, float, complex, rune, byte]")
	}
}

func EqualAny(a, b any) (bool, error) {
	return a == b, nil
}
