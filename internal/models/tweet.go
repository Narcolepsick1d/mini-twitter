package models

type (
	TweetReq struct {
		UserId   string `db:"user_id"`
		Content  string `db:"content" json:"content"`
		MediaUrl string `db:"media_url" json:"media_url"`
	}
)
