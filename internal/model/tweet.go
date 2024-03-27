package model

import (
	"Narcolepsick1d/mini-twitter/internal/models"
	"context"
	"github.com/doug-martin/goqu/v9"
)

func CreateTweet(ctx context.Context, db *goqu.Database, tweetData models.TweetReq) error {
	insertQuery := db.Insert("tweets").
		Cols("user_id", "content", "media_url").
		Vals(goqu.Vals{tweetData.UserId, tweetData.Content, tweetData.MediaUrl})

	_, err := insertQuery.Executor().Exec()
	if err != nil {
		return err
	}
	return nil
}
