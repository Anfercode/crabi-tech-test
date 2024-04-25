package value_objects

import (
	"errors"

	"github.com/google/uuid"
)

type UUIDv4 struct {
	value string
}

func NewUUIDv4(input string) (*UUIDv4, error) {
	id, err := uuid.Parse(input)
	if err != nil {
		return nil, errors.New("input is not a valid UUID")
	}

	if id.Version() != 4 {
		return nil, errors.New("UUID must be version 4")
	}

	return &UUIDv4{value: input}, nil
}

func (u UUIDv4) String() string {
	return u.value
}
