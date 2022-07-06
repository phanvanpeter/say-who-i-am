package data

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

type Validation struct {
	validate *validator.Validate
}

type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("Key: '%s' Error: Field validation for '%s' which failed on '%s'",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Errors() []string {
	var errs []string

	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

// NewValidator creates a new validator
func NewValidator() *Validation {
	validate := validator.New()
	_ = validate.RegisterValidation("phone", validatePhoneNum)

	return &Validation{
		validate: validate,
	}
}

func (v Validation) Validate(i interface{}) ValidationErrors {
	err := v.validate.Struct(i)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)

	if len(errs) == 0 {
		return nil
	}

	var returnErrs []ValidationError
	for _, err := range errs {
		returnErrs = append(returnErrs, ValidationError{err})
	}
	return returnErrs
}

// validatePhoneNum validates a given phone number
func validatePhoneNum(fl validator.FieldLevel) bool {
	// TODO validate phone number in a better way
	reg := regexp.MustCompile(`\d+-\d+-\d+`)
	matches := reg.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}
