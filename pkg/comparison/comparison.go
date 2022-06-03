package comparison

import (
	"fmt"
)

func toFloat(a any) (float64, error) {
	switch a := a.(type) {
	case int:
		return float64(a), nil
	case int8:
		return float64(a), nil
	case int16:
		return float64(a), nil
	case int32:
		return float64(a), nil
	case int64:
		return float64(a), nil
	case float32:
		return float64(a), nil
	case byte:
		return float64(a), nil
	case float64:
		return a, nil
	default:
		return 0.0, fmt.Errorf("non converable type")
	}

}

func GraterAny(a, b any) (bool, error) {
	as, ok := a.(string)
	if ok {
		bs, ok := b.(string)
		if ok {
			return as > bs, nil
		}
	}

	af, err := toFloat(a)
	if err != nil {
		return false, fmt.Errorf("can't convert first argument to float64")
	}

	bf, err := toFloat(b)
	if err != nil {
		return false, fmt.Errorf("can't convert second argument to float64")
	}

	return af > bf, nil
}

func LesserAny(a, b any) (bool, error) {
	as, ok := a.(string)
	if ok {
		bs, ok := b.(string)
		if ok {
			return as < bs, nil
		}
	}

	af, err := toFloat(a)
	if err != nil {
		return false, fmt.Errorf("can't convert first argument to float64")
	}

	bf, err := toFloat(b)
	if err != nil {
		return false, fmt.Errorf("can't convert second argument to float64")
	}

	return af < bf, nil
}

func EqualAny(a, b any) bool {
	af, aerr := toFloat(a)
	bf, berr := toFloat(b)

	if aerr == nil && berr == nil {
		return af == bf
	}

	return a == b
}
