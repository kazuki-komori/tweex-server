package database

import (
	"fmt"
	"time"

	"github.com/kazuki-komori/tweex-server/domain/entity"
	"github.com/kazuki-komori/tweex-server/domain/repository"
)

type Userrepository struct {
	SqlHandler
}

func NewUserRepository(sqlHandler SqlHandler) repository.UserRepository {
	UserRepository := Userrepository{sqlHandler}
	return &UserRepository
}

func (r *Userrepository) AddUser(userEntity entity.User) error {
	db := r.SqlHandler.db
	dto := new(user)
	dto.UID = userEntity.UID
	dto.DisplayName = userEntity.DisplayName
	dto.AccessToken = userEntity.AccessToken
	res := db.Create(&dto)
	if res.Error != nil {
		return fmt.Errorf("failed to create user=%w", res.Error)
	}
	return nil
}

type user struct {
	UID         string    `json:"uid"`
	DisplayName string    `json:"display_name"`
	AccessToken string    `json:"access_token"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
