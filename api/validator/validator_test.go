package validator

import "testing"

func TestValidateFormPassword_ValidForm(t *testing.T) {
	mockForm := make(map[string][]string)
	mockForm["password"] = []string{"test"}
	err := ValidateFormPassword(mockForm)
	if err != nil {
		t.Errorf("ValidateFormPassword was incorrect, got: %v, expected: %v.", err.Error(), nil)
	}
}

func TestValidateFormPassword_MissingPassword(t *testing.T) {
	mockForm := make(map[string][]string)
	err := ValidateFormPassword(mockForm)
	if err != errMissingPassword {
		t.Errorf("ValidateFormPassword was incorrect, got: %v, expected: %v.", err, errMissingPassword.Error())
	}
}

func TestValidateFormPassword_EmptyPassword(t *testing.T) {
	mockForm := make(map[string][]string)
	mockForm["password"] = []string{""}
	err := ValidateFormPassword(mockForm)
	if err != errMissingPassword {
		t.Errorf("ValidateFormPassword was incorrect, got: %v, expected: %v.", err, errMissingPassword.Error())
	}
}
