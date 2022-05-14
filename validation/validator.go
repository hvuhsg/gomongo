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
var validFilterExpressions = []string{
	"$gt",
	"$gte",
	"$eq",
	"$lt",
	"$lte",
	"$ne",
	"$exists",
	"$and",
	"$or",
	"$nor",
	"$not",
	"$in",
	"$nin",
}
var validTopLevelFilters = []string{
	"$not",
	"$and",
	"$or",
	"$nor",
}

// Valid mutations for document updates
var validMutationExpressions = []string{
	"$set",
	"$unset",
	"$rename",
	"$inc",
	"$mul",
	"$min",
	"$max",
	"$push",
	"$pull",
	"$pop",
	"$pullAll",
	"$addToSet",
	"$each",
	"$position",
	"$slice",
	"$sort",
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

func (validator Validator) ValidateMutation(mutation map[string]interface{}) error {
	for expression, value := range mutation {
		isExpressionValid := false

		for _, validExpression := range validMutationExpressions {
			if validExpression == expression {
				isExpressionValid = true
			}
		}

		if !isExpressionValid {
			return fmt.Errorf("invalid mutation expression '%s'", expression)
		}

		switch v := value.(type) {
		case map[string]interface{}:
			err := validator.ValidateMutation(v)

			if err != nil {
				return err
			}
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

func (validator Validator) ValidateFilter(filter map[string]interface{}) error {
	for field, expression := range filter {
		if field[0] == '$' {
			validField := false
			for _, validTopLevelFilter := range validTopLevelFilters {
				if field == validTopLevelFilter {
					validField = true
				}
			}

			if !validField {
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
				isExpressionValid := false

				for _, validExpression := range validFilterExpressions {
					if validExpression == expression {
						isExpressionValid = true
					}
				}

				if !isExpressionValid {
					return fmt.Errorf("invalid filter expression '%s'", expression)
				}
			}
		}
	}

	return nil
}
