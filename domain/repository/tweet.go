package repository

import "github.com/kazuki-komori/tweex-server/domain/entity"

type TweetRepository interface {
	AddTweet(tweet entity.Tweet) error
	GetTweetByUID(uid string) (tweets []entity.Tweet, err error)
}
