package helpers

import (
	"errors"
	"regexp"
	"strconv"
)

// const TagName = "validate"

// type validator struct {
// 	FieldName  string
// 	FieldValue string // value is the input value
// 	IsRequired bool   // definds the wheather the input value required or not
// 	MaxLength  int    // definds maximum length of the input value,  0 value definds no max length check
// 	MinLength  int    // definds minimum length of the input value, 0 value definds no min length check
// 	Regex      string // definds regex of the input value, "" value definds no regex required
// }

// // Validate is used to validate any input data but in string type
// // @returns error if any
// func (v validator) validate() []error {
// 	var errs []error
// 	length := len(v.FieldValue)
// 	Re := regexp.MustCompile(v.Regex)
// 	if v.IsRequired && length < 1 {
// 		errs = append(errs, errors.New(v.FieldName+" is Required"))
// 	}

// 	// Min length check
// 	// If params min length value is zero that indecates, there will be no min length check
// 	if v.MaxLength != 0 && length > 1 && length < v.MinLength {
// 		errs = append(errs, errors.New(v.FieldName+" must be min "+strconv.Itoa(v.MinLength)))
// 	}

// 	// Max length check
// 	// If params max length value is zero that indecates, there will be no max length check
// 	if v.MaxLength != 0 && length > 1 && length > v.MaxLength {
// 		errs = append(errs, errors.New(v.FieldName+" must be max "+strconv.Itoa(v.MaxLength)))
// 	}

// 	if !Re.MatchString(v.FieldValue) { // Regex check
// 		errs = append(errs, errors.New("Invalid "+v.FieldName))
// 	}

// 	return errs
// }

// func Validator(val interface{}) []error {
// 	v := reflect.ValueOf(val)
// 	validatorObj := validator{}
// 	for i := 0; i < v.NumField(); i++ {
// 		v.Type().Field(i).Tag.Get(TagName)
// 		tag := v.Type().Field(i).Tag.Get(TagName)
// 		if tag == "" || tag == "-" {
// 			continue
// 		}
// 		fmt.Sscanf(tag, "required=%b,min=%d,max=%d,regex=%s", &validatorObj.IsRequired, &validatorObj.MinLength, &validatorObj.MaxLength, &validatorObj.Regex)
// 		validatorObj.FieldName = v.Type().Field(i).Name
// 		validatorObj.FieldValue = v.Field(i).String()
// 	}
// 	return validatorObj.validate()
// }

func Validator(value string, isRequired bool, minLength, maxLength int, regex, fieldName string) error {

	length := len(value)
	Re := regexp.MustCompile(regex)
	if isRequired && length < 1 {
		return errors.New(fieldName + " is Required")
	}

	// Min length check
	// If params min length value is zero that indecates, there will be no min length check
	if minLength != 0 && length > 1 && length < minLength {
		return errors.New(fieldName + " must be min " + strconv.Itoa(minLength))
	}

	// Max length check
	// If params max length value is zero that indecates, there will be no max length check
	if maxLength != 0 && length > 1 && length > maxLength {
		return errors.New(fieldName + " must be max " + strconv.Itoa(maxLength))
	}

	if !Re.MatchString(value) { // Regex check
		return errors.New("Invalid " + fieldName)
	}

	return nil
}
