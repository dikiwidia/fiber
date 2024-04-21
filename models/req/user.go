package req

import (
	"database/sql"
	"time"
)

type User struct {
	Name         string         `validate:"required"` // A regular string field
	Email        *string        `validate:"required"` // A pointer to a string, allowing for null values
	Age          uint8          `validate:"required"` // An unsigned 8-bit integer
	Birthday     *time.Time     `validate:"required"` // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
}
