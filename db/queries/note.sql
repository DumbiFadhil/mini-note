-- Create a new note
-- name: CreateNote :one
INSERT INTO notes (title, content, author_id)
VALUES ($1, $2, $3)
RETURNING note_id, title, content, created_at, updated_at;

-- Get all notes for a user
-- name: GetNotesByUser :many
SELECT note_id, title, content, created_at, updated_at
FROM notes
WHERE author_id = $1;

-- Get a single note by ID
-- name: GetNoteByID :one
SELECT note_id, title, content, created_at, updated_at
FROM notes
WHERE note_id = $1 AND author_id = $2;

-- Update a note
-- name: UpdateNote :one
UPDATE notes
SET title = $1, content = $2, updated_at = $3
WHERE note_id = $4 AND author_id = $5
RETURNING note_id, title, content, created_at, updated_at;

-- Delete a note (soft delete)
-- name: DeleteNote :one
UPDATE notes
SET deleted = true, updated_at = $2
WHERE note_id = $1 AND author_id = $3
RETURNING note_id, title, content, created_at, updated_at;

-- Delete a note (hard delete)
-- name: HardDeleteNote :one
DELETE FROM notes
WHERE note_id = $1 AND author_id = $2
RETURNING note_id, title, content, created_at, updated_at;