package domain

type UserRepository interface {
	Save(user *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id string) (*User, error)
}
