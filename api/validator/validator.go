package validator

import "errors"

var errMissingPassword = errors.New("password was not found in form data")

// ValidateFormPassword will throw an error if a password is not included in the form
func ValidateFormPassword(form map[string][]string) error {
	if len(form["password"]) == 0 || len(form["password"][0]) == 0 {
		return errMissingPassword
	}
	return nil
}
