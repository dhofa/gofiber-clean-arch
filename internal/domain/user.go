package domain

import "github.com/dhofa/gofiber-clean-arch/internal/entity"

type UserRepository interface {
	FindAll() ([]entity.User, error)
	FindByID(id uint) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id uint) error
}

type UserUsecase interface {
	GetAll() ([]entity.User, error)
	GetByID(id uint) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id uint) error
}
