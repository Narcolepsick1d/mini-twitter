package model

import (
	"Narcolepsick1d/mini-twitter/internal/models"
	"context"
	"github.com/doug-martin/goqu/v9"
)

func AddLike(ctx context.Context, db *goqu.Database, tweetData models.LikeReq) error {
	insertQuery := db.Insert(models.TableLike).
		Cols("user_id", "tweet_id").
		Vals(goqu.Vals{tweetData.UserId, tweetData.TweetId})

	_, err := insertQuery.Executor().Exec()
	if err != nil {
		return err
	}
	return nil
}
func RemoveLike(ctx context.Context, db *goqu.Database, tweetData goqu.Ex) error {
	_, err := db.Delete(models.TableLike).
		Where(tweetData).
		Executor().Exec()

	if err != nil {
		return err
	}
	return nil
}
