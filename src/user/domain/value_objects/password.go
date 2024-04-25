package value_objects

import (
	"errors"
	"unicode"
)

type Password struct {
	value string
}

func NewPassword(password string) (*Password, error) {
	if len(password) < 8 {
		return nil, errors.New("password must be at least 8 characters long")
	}
	if !hasNumbers(password) {
		return nil, errors.New("password must include at least one number")
	}
	if !hasUpper(password) {
		return nil, errors.New("password must include at least one uppercase letter")
	}
	if !hasSpecialChars(password) {
		return nil, errors.New("password must include at least one special character")
	}
	return &Password{value: password}, nil
}

func hasNumbers(s string) bool {
	for _, r := range s {
		if unicode.IsNumber(r) {
			return true
		}
	}
	return false
}

func hasUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

func hasSpecialChars(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return true
		}
	}
	return false
}

func (p *Password) String() string {
	return p.value
}
