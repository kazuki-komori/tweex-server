package usecase

import (
	"github.com/kazuki-komori/tweex-server/domain/entity"
	"github.com/kazuki-komori/tweex-server/domain/repository"
)

type TweetUsecase struct {
	tweetRepository repository.TweetRepository
}

func NewTweetUsecase(tweetRepo repository.TweetRepository) *TweetUsecase {
	return &TweetUsecase{tweetRepository: tweetRepo}
}

func (u *TweetUsecase) AddTweet(tweet entity.Tweet) (err error) {
	err = u.tweetRepository.AddTweet(tweet)
	return
}

func (u *TweetUsecase) GetTweetByUID(uid string) (tweets []entity.Tweet, err error) {
	tweets, err = u.tweetRepository.GetTweetByUID(uid)
	return
}
