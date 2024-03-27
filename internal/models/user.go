package models

const (
	TableUsers = "users"
)

type (
	User struct {
		Id       string `db:"id"`
		Name     string `db:"name" json:"name"`
		Email    string `db:"email" json:"email"`
		Password string `db:"password" json:"password"`
	}
	UserSignUp struct {
		Name     string `db:"name" json:"name"`
		Email    string `db:"email" json:"email"`
		Password string `db:"password" json:"password"`
	}
)
