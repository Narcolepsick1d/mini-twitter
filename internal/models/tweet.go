package models

const (
	TableTweets   = "tweets"
	TableRetweets = "retweets"
)

type (
	TweetReq struct {
		UserId   string `db:"user_id"`
		Content  string `db:"content" json:"content"`
		MediaUrl string `db:"media_url" json:"media_url"`
	}
	RetweetReq struct {
		TweetId string `db:"tweet_id" json:"tweet_id"`
		UserId  string `db:"user_id"`
	}
)
