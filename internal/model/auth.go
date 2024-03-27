package model

import (
	"Narcolepsick1d/mini-twitter/internal/models"
	"context"
	"github.com/doug-martin/goqu/v9"
	"log"
)

func SignUp(ctx context.Context, db *goqu.Database, user models.UserSignUp) error {
	_, err := db.
		Insert(models.TableUsers).
		Rows(user).
		OnConflict(
			goqu.DoNothing(),
		).Executor().Exec()
	if err != nil {
		return err
	}
	return nil
}
func CreateToken(ctx context.Context, db *goqu.Database, token models.RefreshSession) error {
	_, err := db.
		Insert(models.TokenTable).
		Rows(token).
		OnConflict(
			goqu.DoNothing(),
		).Executor().Exec()
	if err != nil {
		return err
	}
	return nil
}
func GetToken(ctx context.Context, db *goqu.Database, ex *goqu.Ex) error {
	var result models.RefreshSession
	query := db.From(models.TokenTable)
	if ex != nil {
		query = query.Where(*ex)
	}
	err := query.ScanStructs(&result)
	if err != nil {
		return err
	}
	deleteQuery := db.Delete(models.TableUsers).
		Where(goqu.Ex{
			"user_id": result.UserID,
		})
	_, err = deleteQuery.Executor().Exec()
	if err != nil {
		log.Fatal(err)
	}
	return err
}
func GetCredential(ctx context.Context, db *goqu.Database, ex *goqu.Ex) ([]models.User, error) {
	var result []models.User
	query := db.From(models.TableUsers)
	if ex != nil {
		query = query.Where(*ex)
	}

	err := query.ScanStructs(&result)
	if err != nil {
		return result, err
	}
	return result, err
}
