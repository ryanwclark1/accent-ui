package types

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        int
	UserID    uuid.UUID
	Username  string
	Credits   int
	CreatedAt time.Time
}
