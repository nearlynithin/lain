// source: queries.sql

package lain

import (
	"context"
	"time"
)

const createUser = `--name: CreateUser : one
INSERT INTO users (id, email, username)
VALUES ($1, $2, $3)
RETURNING created_at;
`

type CreateUserParams struct {
	UserID   string
	Email    string
	Username string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (time.Time, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.UserID, arg.Email, arg.Username)
	var created_at time.Time
	err := row.Scan(&created_at)
	return created_at, err
}

const userByEmail = `-- name: UserByEmail :one
SELECT id, email, username, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) UserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, userByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const userExistsByEmail = `-- name: UserExistsByEmail :one
SELECT EXISTS (
    SELECT 1 FROM users WHERE email = $1
)
`

func (q *Queries) UserExistsByEmail(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, userExistsByEmail, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const userExistsByUsername = `-- name: UserExistsByUsername :one
SELECT EXISTS (
    SELECT 1 FROM users WHERE username ILIKE $1
)
`

func (q *Queries) UserExistsByUsername(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRowContext(ctx, userExistsByUsername, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
