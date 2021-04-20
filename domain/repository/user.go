package repository

import "github.com/kazuki-komori/tweex-server/domain/entity"

type UserRepository interface {
	AddUser(user entity.User) error
}
