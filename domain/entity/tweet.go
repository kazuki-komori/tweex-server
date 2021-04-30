package entity

type Tweet struct {
	TweetID int64  `json:"tweet_id"`
	UserID  string `json:"user_id"`
	TagID   string `json:"tag_id"`
}
