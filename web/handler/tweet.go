package handler

import (
	"net/http"

	"github.com/kazuki-komori/tweex-server/domain/entity"
	"github.com/kazuki-komori/tweex-server/usecase"
	"github.com/labstack/echo"
)

type TweetHandler struct {
	TweetUC *usecase.TweetUsecase
}

func NewTweetHandler(tweetUC *usecase.TweetUsecase) *TweetHandler {
	return &TweetHandler{TweetUC: tweetUC}
}

// POST /tweet/register tweet を登録
func (h *TweetHandler) PostTweet(c echo.Context) (err error) {
	tweet := new(entity.Tweet)
	c.Bind(tweet)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid params")
	}
	err = h.TweetUC.AddTweet(*tweet)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to register Tweet")
	}
	return c.JSON(http.StatusCreated, tweet)
}

//GET /tweet/all ユーザー登録したツイートを取得
func (h *TweetHandler) GetTweetByID(c echo.Context) error {
	uid := c.Param("uid")
	tweets, err := h.TweetUC.GetTweetByUID(uid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get Tweets")
	}
	return c.JSON(http.StatusCreated, tweets)
}
