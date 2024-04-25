package persistence

import (
	"crabi-tech-test/src/user/domain"
	"errors"
	"sync"
)

type InMemoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]*domain.User
	ids   map[string]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*domain.User),
		ids:   make(map[string]*domain.User),
	}
}

func (repo *InMemoryUserRepository) Save(user *domain.User) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.users[user.Email]; exists {
		return errors.New("user already exists")
	}
	repo.users[user.Email] = user
	repo.ids[user.ID] = user
	return nil
}

func (repo *InMemoryUserRepository) FindByEmail(email string) (*domain.User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if user, exists := repo.users[email]; exists {
		return user, nil
	}
	return nil, errors.New("user not found")
}

func (repo *InMemoryUserRepository) FindByID(id string) (*domain.User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if user, exists := repo.ids[id]; exists {
		return user, nil
	}
	return nil, errors.New("user not found")
}
