package filtering

import (
	"fmt"

	comparison "github.com/hvuhsg/gomongo/pkg/comparison"
)

func Filter(filter_ map[string]interface{}, document map[string]interface{}) (bool, error) {
	for topLevel, expression := range filter_ {
		if topLevel[0] == '$' {
			expression, ok := expression.(map[string]interface{})
			if !ok {
				return false, fmt.Errorf("expression of top level filter must be of type 'map[string]interface{}'")
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
			continue
		}

		expressionMap, ok := expression.(map[string]interface{})
		if !ok {
			expressionMap = map[string]any{"$eq": expression}
		}

		documentValue, ok := document[topLevel]
		for operator, value := range expressionMap {
			switch operator {
			case "$eq":
				if ok && !comparison.EqualAny(value, documentValue) {
					return false, nil
				}
			case "$ne":
				if ok && comparison.EqualAny(value, documentValue) {
					return false, nil
				}
			case "$exists":
				value, isBool := value.(bool)
				if !isBool {
					return false, fmt.Errorf("$exists must have bool type")
				}
				if (ok && !value) || (!ok && value) {
					return false, nil
				}
			case "$gt":
				isGrater, err := comparison.GraterAny(documentValue, value)
				if err != nil || !isGrater {
					return false, err
				}
			case "$gte":
				if ok && !comparison.EqualAny(value, documentValue) {
					isGrater, err := comparison.GraterAny(documentValue, value)
					if err != nil || !isGrater {
						return false, err
					}
				}
			case "$lt":
				isLesser, err := comparison.LesserAny(documentValue, value)
				if err != nil || !isLesser {
					return false, err
				}
			case "$lte":
				if ok && !comparison.EqualAny(value, documentValue) {
					isLesser, err := comparison.LesserAny(documentValue, value)
					if err != nil || !isLesser {
						return false, err
					}
				}
			}
		}
	}
	return true, nil
}

func FilterList(filter map[string]any, list []map[string]any) []map[string]any {
	filteredList := make([]map[string]any, 0, len(list))
	for _, item := range list {
		result, err := Filter(filter, item)
		if err == nil && result {
			filteredList = append(filteredList, item)
		}
	}

	return filteredList
}
