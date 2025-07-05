package usecase

import (
	"github.com/dhofa/gofiber-clean-arch/internal/domain"
	"github.com/dhofa/gofiber-clean-arch/internal/entity"
)

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(r domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{r}
}

func (u *UserUsecase) GetAll() ([]entity.User, error) {
	return u.repo.FindAll()
}

func (u *UserUsecase) GetByID(id uint) (*entity.User, error) {
	return u.repo.FindByID(id)
}

func (u *UserUsecase) Create(user *entity.User) error {
	return u.repo.Create(user)
}

func (u *UserUsecase) Update(user *entity.User) error {
	return u.repo.Update(user)
}

func (u *UserUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
