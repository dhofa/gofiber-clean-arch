package usecase

import (
	"github.com/dhofa/gofiber-clean-arch/internal/domain"
	"github.com/dhofa/gofiber-clean-arch/internal/entity"
)

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(r domain.UserRepository) domain.UserUsecase {
	return &userUsecase{r}
}

func (u *userUsecase) GetAll() ([]entity.User, error) {
	return u.repo.FindAll()
}

func (u *userUsecase) GetByID(id uint) (*entity.User, error) {
	return u.repo.FindByID(id)
}

func (u *userUsecase) Create(user *entity.User) error {
	return u.repo.Create(user)
}

func (u *userUsecase) Update(user *entity.User) error {
	return u.repo.Update(user)
}

func (u *userUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
