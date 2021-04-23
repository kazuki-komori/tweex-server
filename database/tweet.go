package database

import (
	"fmt"

	"github.com/kazuki-komori/tweex-server/domain/entity"
	"github.com/kazuki-komori/tweex-server/domain/repository"
)

type TweetRepository struct {
	SqlHandler
}

func NewTweetrRepository(sqlHandler SqlHandler) repository.TweetRepository {
	TweetRepository := TweetRepository{sqlHandler}
	return &TweetRepository
}

func (r *TweetRepository) AddTweet(tweetEntity entity.Tweet) error {
	db := r.SqlHandler.db
	dto := new(tweet)
	dto.TID = tweetEntity.TID
	dto.UserID = tweetEntity.UserID
	res := db.Create(&dto)
	fmt.Println(res.Error)
	if res.Error != nil {
		return fmt.Errorf("failed to create tweet=%w", res.Error)
	}
	return nil
}

func (r *TweetRepository) GetTweetByUID(uid string) (tweets []entity.Tweet, err error) {
	db := r.SqlHandler.db
	db.Find(&tweets)
	return tweets, err
}

type tweet struct {
	TID    int64  `json:"t_id"`
	UserID string `json:"user_id"`
}
