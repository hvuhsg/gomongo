package engine

import (
	"fmt"

	comparison "github.com/hvuhsg/gomongo/pkg/comparison"
)

/// UTILS START
// Valid filter expressions
var validFilterOperations = map[string]struct{}{
	"$gt":     {},
	"$gte":    {},
	"$eq":     {},
	"$lt":     {},
	"$lte":    {},
	"$ne":     {},
	"$exists": {},
	"$and":    {},
	"$or":     {},
	"$nor":    {},
	"$not":    {},
	"$in":     {},
	"$nin":    {},
}
var validTopLevelFilters = map[string]struct{}{
	"$not": {},
	"$and": {},
	"$or":  {},
	"$nor": {},
}

func IsTopLevelFilterOp(filterOp string) bool {
	_, ok := validTopLevelFilters[filterOp]
	return ok
}

func IsFilterOp(filterOp string) bool {
	_, ok := validTopLevelFilters[filterOp]
	if ok {
		return ok
	}

	_, ok = validFilterOperations[filterOp]
	return ok
}

func ToFilter(filter any) (map[string]any, bool) {
	mapFilter, ok := filter.(map[string]any)
	return mapFilter, ok
}

/// UTILS END

func isContains(collection []any, item any) bool {
	for _, i := range collection {
		isEqual := comparison.EqualAny(i, item)
		if isEqual {
			return true
		}
	}
	return false
}

func eq(documentValue, filterValue any) bool {
	return comparison.EqualAny(filterValue, documentValue)
}

func ne(documentValue, filterValue any) bool {
	return !comparison.EqualAny(filterValue, documentValue)
}

func exists(isExists bool, needToExists any) (bool, error) {
	needToExistsBool, ok := needToExists.(bool)
	if !ok {
		return false, fmt.Errorf("$exists must have bool type")
	}
	if (isExists && !needToExistsBool) || (!isExists && needToExistsBool) {
		return false, nil
	}
	return true, nil
}

func gt(documentValue, filterValue any) (bool, error) {
	isGrater, err := comparison.GraterAny(documentValue, filterValue)
	return isGrater, err
}

func gte(documentValue, filterValue any) (bool, error) {
	isGrater, err := gt(documentValue, filterValue)
	if isGrater || err != nil {
		return isGrater, err
	}

	return eq(documentValue, filterValue), nil
}

func lt(documentValue, filterValue any) (bool, error) {
	isLesser, err := comparison.LesserAny(documentValue, filterValue)
	return isLesser, err
}

func lte(documentValue, filterValue any) (bool, error) {
	isLesser, err := lt(documentValue, filterValue)
	if isLesser || err != nil {
		return isLesser, err
	}

	return eq(documentValue, filterValue), nil
}

func in(documentValue, filterValue any) bool {
	switch collection := documentValue.(type) {
	case []any:
		return isContains(collection, filterValue)
	default:
		return false
	}
}

func nin(documentValue, filterValue any) bool {
	switch collection := documentValue.(type) {
	case []any:
		return !isContains(collection, filterValue)
	default:
		return false
	}
}

func Filter(filter_ map[string]any, document map[string]any) (bool, error) {
	for key, value := range filter_ {
		if IsTopLevelFilterOp(key) {
			expression, ok := ToFilter(value)
			if !ok {
				return false, fmt.Errorf("expression of top level filter must be of type 'map[string]any'")
			}

			switch key {
			case "$and":
				for subfilter, subexpression := range expression {
					result, err := Filter(map[string]any{subfilter: subexpression}, document)

					if err != nil || !result {
						return false, err
					}
				}
			case "$or":
				var valid bool = false

				for subfilter, subexpression := range expression {
					result, err := Filter(map[string]any{subfilter: subexpression}, document)

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
					result, err := Filter(map[string]any{subfilter: subexpression}, document)
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
					result, err := Filter(map[string]any{subfilter: subexpression}, document)
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

		expressionMap, ok := ToFilter(value)
		if !ok {
			expressionMap = map[string]any{"$eq": value}
		}

		documentValue, fieldExists := document[key]
		for operator, filterValue := range expressionMap {
			switch operator {
			case "$eq":
				if fieldExists && !eq(documentValue, filterValue) {
					return false, nil
				}
			case "$ne":
				if fieldExists && !ne(documentValue, filterValue) {
					return false, nil
				}
			case "$exists":
				ok, err := exists(fieldExists, filterValue)
				if err != nil || !ok {
					return false, err
				}
			case "$gt":
				ok, err := gt(documentValue, filterValue)
				if err != nil || !ok {
					return false, err
				}
			case "$gte":
				ok, err := gte(documentValue, filterValue)
				if err != nil || !ok {
					return false, err
				}
			case "$lt":
				ok, err := lt(documentValue, filterValue)
				if err != nil || !ok {
					return false, err
				}
			case "$lte":
				ok, err := lte(documentValue, filterValue)
				if err != nil || !ok {
					return false, err
				}
			case "$in":
				if !in(documentValue, filterValue) {
					return false, nil
				}
			case "$nin":
				if !nin(documentValue, filterValue) {
					return false, nil
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
