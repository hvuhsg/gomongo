package validation_test

import (
	"fmt"
	"testing"

	"github.com/hvuhsg/gomongo/engine/validation"
)

func testNameValidity(name string) bool {
	validator := validation.New()
	err := validator.ValidateName(name)
	return err == nil
}

func testFilterValidity(filter map[string]interface{}) bool {
	validator := validation.New()
	err := validator.ValidateFilter(filter)
	return err == nil
}

func testMutationValidity(filter map[string]interface{}) bool {
	validator := validation.New()
	err := validator.ValidateMutation(filter)
	if err != nil {
		fmt.Println(err)
	}
	return err == nil
}

func TestNameValidation(t *testing.T) {
	invalid_names := []string{"1hello", "Capital", "spacial%"}
	valid_names := []string{"hello", "snake_case", "number5", "_id"}

	for _, invalid_name := range invalid_names {
		t.Run(invalid_name, func(t *testing.T) {
			is_valid := testNameValidity(invalid_name)
			if is_valid {
				t.Errorf("'%s' sholud be invalid name", invalid_name)
			}
		})
	}

	for _, valid_name := range valid_names {
		t.Run(valid_name, func(t *testing.T) {
			is_valid := testNameValidity(valid_name)
			if !is_valid {
				t.Errorf("'%s' sholud be valid name", valid_name)
			}
		})
	}
}

func TestDocumentValidation(t *testing.T) {
	validator := validation.New()
	document := map[string]interface{}{
		"hello":     "string",
		"number":    5,
		"float":     5.4,
		"bool":      true,
		"arr":       []int{5, 8, -9},
		"validator": validator,
	}

	err := validator.ValidateDocument(document)
	if err != nil {
		t.Error(err)
	}
}

func TestFilterValidation(t *testing.T) {
	invalidFilters := []map[string]interface{}{
		{"$inc": "value"},
		{"$gt": map[string]interface{}{"k": "v"}},
		{"$$food": "ok"},
		{"key": map[string]interface{}{"$set": "b"}},
		{"$set": map[string]interface{}{"a": "b"}},
	}
	validFilters := []map[string]interface{}{
		{"key": "value"},
		{"key": "$value"},
		{"key": map[string]interface{}{"$gt": 5}},
		{"key": map[string]interface{}{"$exists": true}},
	}

	for _, invalidFilter := range invalidFilters {
		t.Run("test invalid filter", func(t *testing.T) {
			isValid := testFilterValidity(invalidFilter)
			if isValid {
				t.Errorf("'%v' sholud be invalid filter", invalidFilter)
			}
		})
	}

	for _, validFilter := range validFilters {
		t.Run("test valid filter", func(t *testing.T) {
			isValid := testFilterValidity(validFilter)
			if !isValid {
				t.Errorf("'%v' sholud be valid filter", validFilter)
			}
		})
	}
}

func TestMutationValidation(t *testing.T) {
	invalidMutations := []map[string]interface{}{
		{"$inc": "value"},
		{"$gt": 5},
		{"$mul": map[string]interface{}{"k": "v"}},
		{"key": "value"},
		{"a": map[string]interface{}{"$set": "b"}},
	}
	validMutations := []map[string]interface{}{
		{"$unset": "value"},
		{"$set": map[string]interface{}{"key": "value"}},
		{"$inc": map[string]interface{}{"age": 1}},
		{"$mul": map[string]interface{}{"age": 1}},
		{"$push": map[string]interface{}{"arr": 5}},
		{"$inc": map[string]interface{}{"age": 1, "count": 2}, "$set": map[string]interface{}{"name": "mosh"}},
	}

	for _, invalidMutation := range invalidMutations {
		t.Run("test invalid mutation", func(t *testing.T) {
			isValid := testMutationValidity(invalidMutation)
			if isValid {
				t.Errorf("'%v' sholud be invalid mutation", invalidMutation)
			}
		})
	}

	for _, validMutation := range validMutations {
		t.Run("test valid mutation", func(t *testing.T) {
			isValid := testMutationValidity(validMutation)
			if !isValid {
				t.Errorf("'%v' sholud be valid mutation", validMutation)
			}
		})
	}
}
