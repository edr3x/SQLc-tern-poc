// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package query

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        pgtype.UUID
	FirstName pgtype.Text
	LastName  pgtype.Text
	Email     pgtype.Text
	Password  pgtype.Text
}
