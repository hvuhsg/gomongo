package validation

import (
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
	return nil
}

func (validator fakeValidator) ValidateFilter(filter map[string]interface{}) error {
	return nil
}
