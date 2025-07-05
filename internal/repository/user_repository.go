package repository

import (
	"github.com/dhofa/gofiber-clean-arch/internal/domain"
	"github.com/dhofa/gofiber-clean-arch/internal/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) FindByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}
