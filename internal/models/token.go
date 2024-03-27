package models

import (
	"time"
)

const TokenTable = "refresh_tokens"

type RefreshSession struct {
	UserID    string    `db:"user_id" json:"user_id"`
	Token     string    `db:"token" json:"token"`
	ExpiresAt time.Time `db:"expires_at" json:"expires_at"`
}
