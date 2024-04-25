package persistence

import (
	"crabi-tech-test/src/user/domain"
	"errors"

	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Save(user *domain.User) error {
	result := r.db.Save(user)
	return result.Error
}

func (r *PostgresUserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *PostgresUserRepository) FindByID(id string) (*domain.User, error) {
	var user domain.User
	result := r.db.First(&user, id)
	return &user, result.Error
}
