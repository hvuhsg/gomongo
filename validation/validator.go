package validation

import (
	"errors"
	"strings"

	"github.com/hvuhsg/gomongo/engine"
)

type fakeValidator struct {
}

func NewFakeValidator() engine.IValidator {
	validator := new(fakeValidator)
	return validator
}

func (validator fakeValidator) ValidateDocument(docuemnt map[string]interface{}) error {
	return nil
}

func (validator fakeValidator) ValidateMutation(mutation map[string]interface{}) error {
	return nil
}

func (validator fakeValidator) ValidateName(name string) error {
	invalid_chars := "!@#$%^&*() "

	for _, invalid_chr := range invalid_chars {
		if strings.Contains(name, string(invalid_chr)) {
			return errors.New("invalid character in name")
		}
	}

	first_chr := name[0]
	if '0' < first_chr && '9' > first_chr {
		return errors.New("name can't start with a number")
	}

	return nil
}

func (validator fakeValidator) ValidateFilter(filter map[string]interface{}) error {
	return nil
}
