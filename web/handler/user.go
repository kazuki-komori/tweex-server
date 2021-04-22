package handler

import (
	"net/http"

	"github.com/kazuki-komori/tweex-server/domain/entity"
	"github.com/kazuki-komori/tweex-server/usecase"
	"github.com/labstack/echo"
)

type UserHandler struct {
	UserUC *usecase.UserUsecase
}

func NewUserHandler(userUC *usecase.UserUsecase) *UserHandler {
	return &UserHandler{UserUC: userUC}
}

// POST /user/register ユーザーの登録
func (h *UserHandler) PostUser(c echo.Context) (err error) {
	user := new(entity.User)
	err = c.Bind(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid params")
	}
	err = h.UserUC.AddUser(*user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to register User")
	}
	return c.JSON(http.StatusCreated, user)
}
