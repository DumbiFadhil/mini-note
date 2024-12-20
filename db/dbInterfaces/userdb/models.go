// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package userdb

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Note struct {
	NoteID    int32
	Title     string
	Content   string
	AuthorID  int32
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	Deleted   pgtype.Bool
}

type User struct {
	UserID    int32
	Username  string
	Email     string
	Password  string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
