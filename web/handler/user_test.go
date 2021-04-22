package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kazuki-komori/tweex-server/domain/entity"
	"github.com/kazuki-komori/tweex-server/domain/mock_repository"
	"github.com/kazuki-komori/tweex-server/usecase"
	"github.com/kazuki-komori/tweex-server/web/handler"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T, req *http.Request, rec *httptest.ResponseRecorder) echo.Context {
	// Setup Echo
	e := echo.New()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := e.NewContext(req, rec)
	return c
}

func TestAddUser(t *testing.T) {
	expectJSON := `{"uid": "test uid", "display_name": "name", "access_token": "token"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/", strings.NewReader(expectJSON))
	rec := httptest.NewRecorder()

	c := setup(t, req, rec)
	c.SetPath("/user/register")
	c.SetParamNames("uid", "display_name", "access_token")
	c.SetParamValues("test uid", "name", "token")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_repository.NewMockUserRepository(ctrl)
	uc := usecase.NewUserUsecase(mockRepo)
	h := &handler.UserHandler{UserUC: uc}

	user := new(entity.User)
	user.AccessToken = "token"
	user.UID = "test uid"
	user.DisplayName = "name"
	mockRepo.EXPECT().AddUser(*user)

	err := h.PostUser(c)
	if err != nil {
		t.Errorf("failed to post user = %d", err)
	}

	assert.JSONEq(t, strings.TrimRight(rec.Body.String(), "\n"), expectJSON)
}
