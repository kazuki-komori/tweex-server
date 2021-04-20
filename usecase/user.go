package usecase

import (
	"github.com/kazuki-komori/tweex-server/domain/entity"
	"github.com/kazuki-komori/tweex-server/domain/repository"
)

type UserUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepository: userRepo}
}

func (u *UserUsecase) AddUser(user entity.User) (err error) {
	err = u.AddUser(user)
	return
}
