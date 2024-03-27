package models

const TableLike = "likes"

type (
	LikeReq struct {
		TweetId string `db:"tweet_id" json:"tweet_id"`
		UserId  string `db:"user_id"`
	}
)
