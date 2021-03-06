package web

import (
	"net/http"
	"os"

	"github.com/kazuki-komori/tweex-server/usecase"
	"github.com/kazuki-komori/tweex-server/web/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewServer(userUC *usecase.UserUsecase, tweetUC *usecase.TweetUsecase) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userHandler := handler.NewUserHandler(userUC)
	TweetHandler := handler.NewTweetHandler(tweetUC)

	v1 := e.Group("/api/v1")

	v1.GET("/health", health)
	v1.POST("/user/register", userHandler.PostUser)
	v1.POST("/tweet/register", TweetHandler.PostTweet)
	v1.POST("/tweet/all", TweetHandler.GetTweetByID)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func health(c echo.Context) error {
	type health struct {
		Status string `json:"status"`
	}
	res := new(health)
	res.Status = "OK"
	return c.JSON(http.StatusOK, res)
}
