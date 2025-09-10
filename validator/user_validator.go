package validator

import (
	"CRUD_operation/models"
	"errors"
)

func ValidateUser(user models.Schema) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.College == "" {
		return errors.New("college is required")
	}
	if user.Branch == "" {
		return errors.New("branch is required")
	}
	return nil
}
