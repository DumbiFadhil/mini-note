// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: note.sql

package notedb

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createNote = `-- name: CreateNote :one
INSERT INTO notes (title, content, author_id)
VALUES ($1, $2, $3)
RETURNING note_id, title, content, created_at, updated_at
`

type CreateNoteParams struct {
	Title    string
	Content  string
	AuthorID int32
}

type CreateNoteRow struct {
	NoteID    int32
	Title     string
	Content   string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

// Create a new note
func (q *Queries) CreateNote(ctx context.Context, arg CreateNoteParams) (CreateNoteRow, error) {
	row := q.db.QueryRow(ctx, createNote, arg.Title, arg.Content, arg.AuthorID)
	var i CreateNoteRow
	err := row.Scan(
		&i.NoteID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteNote = `-- name: DeleteNote :one
UPDATE notes
SET deleted = true, updated_at = $2
WHERE note_id = $1 AND author_id = $3
RETURNING note_id, title, content, created_at, updated_at
`

type DeleteNoteParams struct {
	NoteID    int32
	UpdatedAt pgtype.Timestamp
	AuthorID  int32
}

type DeleteNoteRow struct {
	NoteID    int32
	Title     string
	Content   string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

// Delete a note (soft delete)
func (q *Queries) DeleteNote(ctx context.Context, arg DeleteNoteParams) (DeleteNoteRow, error) {
	row := q.db.QueryRow(ctx, deleteNote, arg.NoteID, arg.UpdatedAt, arg.AuthorID)
	var i DeleteNoteRow
	err := row.Scan(
		&i.NoteID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getNoteByID = `-- name: GetNoteByID :one
SELECT note_id, title, content, created_at, updated_at
FROM notes
WHERE note_id = $1 AND author_id = $2
`

type GetNoteByIDParams struct {
	NoteID   int32
	AuthorID int32
}

type GetNoteByIDRow struct {
	NoteID    int32
	Title     string
	Content   string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

// Get a single note by ID
func (q *Queries) GetNoteByID(ctx context.Context, arg GetNoteByIDParams) (GetNoteByIDRow, error) {
	row := q.db.QueryRow(ctx, getNoteByID, arg.NoteID, arg.AuthorID)
	var i GetNoteByIDRow
	err := row.Scan(
		&i.NoteID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getNotesByUser = `-- name: GetNotesByUser :many
SELECT note_id, title, content, created_at, updated_at
FROM notes
WHERE author_id = $1
`

type GetNotesByUserRow struct {
	NoteID    int32
	Title     string
	Content   string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

// Get all notes for a user
func (q *Queries) GetNotesByUser(ctx context.Context, authorID int32) ([]GetNotesByUserRow, error) {
	rows, err := q.db.Query(ctx, getNotesByUser, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetNotesByUserRow
	for rows.Next() {
		var i GetNotesByUserRow
		if err := rows.Scan(
			&i.NoteID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const hardDeleteNote = `-- name: HardDeleteNote :one
DELETE FROM notes
WHERE note_id = $1 AND author_id = $2
RETURNING note_id, title, content, created_at, updated_at
`

type HardDeleteNoteParams struct {
	NoteID   int32
	AuthorID int32
}

type HardDeleteNoteRow struct {
	NoteID    int32
	Title     string
	Content   string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

// Delete a note (hard delete)
func (q *Queries) HardDeleteNote(ctx context.Context, arg HardDeleteNoteParams) (HardDeleteNoteRow, error) {
	row := q.db.QueryRow(ctx, hardDeleteNote, arg.NoteID, arg.AuthorID)
	var i HardDeleteNoteRow
	err := row.Scan(
		&i.NoteID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateNote = `-- name: UpdateNote :one
UPDATE notes
SET title = $1, content = $2, updated_at = $3
WHERE note_id = $4 AND author_id = $5
RETURNING note_id, title, content, created_at, updated_at
`

type UpdateNoteParams struct {
	Title     string
	Content   string
	UpdatedAt pgtype.Timestamp
	NoteID    int32
	AuthorID  int32
}

type UpdateNoteRow struct {
	NoteID    int32
	Title     string
	Content   string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

// Update a note
func (q *Queries) UpdateNote(ctx context.Context, arg UpdateNoteParams) (UpdateNoteRow, error) {
	row := q.db.QueryRow(ctx, updateNote,
		arg.Title,
		arg.Content,
		arg.UpdatedAt,
		arg.NoteID,
		arg.AuthorID,
	)
	var i UpdateNoteRow
	err := row.Scan(
		&i.NoteID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
