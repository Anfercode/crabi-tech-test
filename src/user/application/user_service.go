package application

import (
	"crabi-tech-test/src/user/domain"
	"errors"
)

type UserService struct {
	UserRepository     domain.UserRepository
	BlacklistValidator domain.BlacklistValidator
}

func NewUserService(repo domain.UserRepository, blacklistValidator domain.BlacklistValidator) *UserService {
	return &UserService{
		UserRepository:     repo,
		BlacklistValidator: blacklistValidator,
	}
}

func (s *UserService) CreateUser(id, firstName, lastName, email, password string) (*domain.User, error) {
	blacklisted, err := s.BlacklistValidator.Validate(firstName, lastName, email)
	if err != nil {
		return nil, err
	}

	if blacklisted {
		return nil, errors.New("user is blacklisted")
	}

	existingUser, err := s.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("a user with the same email already exists")
	}

	user, err := domain.NewUser(id, firstName, lastName, email, password)
	if err != nil {
		return nil, err
	}

	err = s.UserRepository.Save(user)
	return user, err
}

func (s *UserService) Login(email, password string) (*domain.User, error) {
	user, err := s.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func (s *UserService) GetUser(email string) (*domain.User, error) {
	return s.UserRepository.FindByEmail(email)
}
