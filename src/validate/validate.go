package validate

import (
	"errors"
	"unicode/utf8"
)

func ValidateName(username string) (correct bool, err error) {
	length := utf8.RuneCountInString(username)
	if length == 0 {
		return false, errors.New("username is required")
	}
	return true, nil
}

func ValidateTitle(username string) (correct bool, err error) {
	length := utf8.RuneCountInString(username)
	if length == 0 {
		return false, errors.New("title is required")
	}
	if length < 2 || length > 20 {
		return false, errors.New("The title field must be between 2-20 chars")
	}
	return true, nil
}
