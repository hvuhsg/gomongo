package filtering

import (
	"fmt"
)

func Filter(filter_ map[string]interface{}, document map[string]interface{}) (bool, error) {
	for topLevel, expression := range filter_ {
		if topLevel[0] == '$' {
			expression, ok := expression.(map[string]map[string]interface{})
			if !ok {
				return false, fmt.Errorf("expression of top level filter must be of type 'map[string]map[string]interface{}'")
			}

			switch topLevel {
			case "$and":
				for subfilter, subexpression := range expression {
					result, err := Filter(map[string]interface{}{subfilter: subexpression}, document)

					if err != nil || !result {
						return false, err
					}
				}
			case "$or":
				var valid bool = false

				for subfilter, subexpression := range expression {
					result, err := Filter(map[string]interface{}{subfilter: subexpression}, document)

					if err != nil {
						return false, err
					}

					valid = result || valid
				}

				if !valid {
					return false, nil
				}
			case "$nor":
				var valid bool = false
				for subfilter, subexpression := range expression {
					result, err := Filter(map[string]interface{}{subfilter: subexpression}, document)
					if err != nil {
						return false, err
					}
					valid = result || valid
				}

				if valid {
					return false, nil
				}
			case "$not":
				if len(expression) != 1 {
					return false, fmt.Errorf("length of $not expression must be 1")
				}

				for subfilter, subexpression := range expression {
					result, err := Filter(map[string]interface{}{subfilter: subexpression}, document)
					if err != nil {
						return false, err
					}

					if result {
						return false, nil
					}
				}
			default:
				return false, fmt.Errorf("top level filter that start with '$' must be one of [$not, $nor, $or, $and]")
			}
		}
	}
	return true, nil
}
