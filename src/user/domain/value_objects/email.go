package value_objects

import (
	"fmt"
	"regexp"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	if !isValidEmail(value) {
		return nil, fmt.Errorf("invalid email address: %s", value)
	}
	return &Email{value: value}, nil
}

func isValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return regex.MatchString(email)
}

func (e Email) String() string {
	return e.value
}
