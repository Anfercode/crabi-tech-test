package domain

type BlacklistValidator interface {
	Validate(firstName string, lastName string, email string) (bool, error)
}
