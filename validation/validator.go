package validation

import (
	"fmt"
	"strings"

	"github.com/hvuhsg/gomongo/engine"
)

// Valid chars in document keys
var validKeyChars = "0123456789abcdefghijklmnopqrstuvwxyz_ "

// Chars that invalid as a first document key char
var invalidFirstChars = "0123456789 "

// Valid filter expressions
var validFilterExpressions = map[string]struct{}{
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

// Valid mutations for document updates
var validMutationOprators = map[string]struct{}{
	"$set":      {},
	"$unset":    {},
	"$rename":   {},
	"$inc":      {},
	"$mul":      {},
	"$min":      {},
	"$max":      {},
	"$push":     {},
	"$pull":     {},
	"$pop":      {},
	"$pullAll":  {},
	"$addToSet": {},
	"$each":     {},
	"$position": {},
	"$slice":    {},
	"$sort":     {},
}

type Validator struct{}

func New() engine.IValidator {
	validator := new(Validator)
	return validator
}

func (validator Validator) ValidateDocument(docuemnt map[string]interface{}) error {
	for key := range docuemnt {
		err := validator.ValidateName(key)
		if err != nil {
			return err
		}
	}

	return nil
}

func (validator Validator) ValidateName(name string) error {
	for _, char := range name {
		if !strings.Contains(validKeyChars, string(char)) {
			return fmt.Errorf("invalid char '%c' in name", char)
		}
	}

	if strings.Contains(invalidFirstChars, string(name[0])) {
		return fmt.Errorf("name can't start with '%c'", name[0])
	}

	return nil
}

func (validator Validator) ValidateMutation(mutation map[string]interface{}) error {
	for operator, expression := range mutation {
		_, isValid := validMutationOprators[operator]

		if !isValid {
			return fmt.Errorf("invalid mutation expression '%s'", operator)
		}

		if operator == "$inc" || operator == "$mul" {
			switch expression := expression.(type) {
			case map[string]int:
			case map[string]interface{}:
				for _, v := range expression {
					switch v.(type) {
					case int:
					default:
						return fmt.Errorf("value for $inc or $mul must be of type map[string]int")
					}
				}
			default:
				return fmt.Errorf("value for $inc or $mul must be of type map[string]int")
			}
		}
	}

	return nil
}

func (validator Validator) ValidateFilter(filter map[string]interface{}) error {
	for field, expression := range filter {
		if field[0] == '$' {
			_, ok := validTopLevelFilters[field]

			if !ok {
				return fmt.Errorf("invalid top level filter '%s'", field)
			}
		} else {
			err := validator.ValidateName(field)
			if err != nil {
				return err
			}
		}

		switch v := expression.(type) {
		case map[string]interface{}:
			for expression := range v {
				_, isValid := validFilterExpressions[expression]

				if !isValid {
					return fmt.Errorf("invalid filter expression '%s'", expression)
				}
			}
		}
	}

	return nil
}
