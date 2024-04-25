package domain

import (
	"crabi-tech-test/src/user/domain/value_objects"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string
	Name     string
	LastName string
	Email    string
	Password string
}

func NewUser(id, name, lastName, email, password string) (*User, error) {
	e, err := value_objects.NewEmail(email)
	if err != nil {
		return nil, err
	}

	uuid, err := value_objects.NewUUIDv4(id)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       uuid.String(),
		Name:     name,
		LastName: lastName,
		Email:    e.String(),
		Password: password,
	}, nil
}

func (u *User) UpdateEmail(newEmail string) error {
	e, err := value_objects.NewEmail(newEmail)
	if err != nil {
		return err
	}
	u.Email = e.String()
	return nil
}
