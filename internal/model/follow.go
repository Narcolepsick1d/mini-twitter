package model

import (
	"Narcolepsick1d/mini-twitter/internal/models"
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
)

func FollowUser(ctx context.Context, db *goqu.Database, follow models.Follow) error {

	rows, err := db.
		Insert(models.TableFollow).
		Cols("lead_id", "follower").
		Vals(goqu.Vals{follow.LeadId, follow.FollowerId}).
		OnConflict(
			goqu.DoNothing(),
		).Executor().Exec()
	if err != nil {
		return err
	}
	a, err := rows.RowsAffected()
	if a == 0 {
		return fmt.Errorf("u already following")
	}
	return nil
}
func UnFollowUser(ctx context.Context, db *goqu.Database, follow goqu.Ex) error {
	_, err := db.
		Delete(models.TableFollow).
		Where(follow).
		Executor().Exec()
	if err != nil {
		return err
	}
	return nil
}
