// Code generated by sqlc. DO NOT EDIT.
// source: users_pass.sql

package database

import (
	"context"
)

const getUsersPass = `-- name: GetUsersPass :one
SELECT id, password
FROM users_pass
WHERE id = $1
`

func (q *Queries) GetUsersPass(ctx context.Context, id int32) (UsersPass, error) {
	row := q.db.QueryRowContext(ctx, getUsersPass, id)
	var i UsersPass
	err := row.Scan(&i.ID, &i.Password)
	return i, err
}

const setUsersPass = `-- name: SetUsersPass :exec
INSERT INTO users_pass (id, password)
VALUES ($1, $2)
`

type SetUsersPassParams struct {
	ID       int32
	Password string
}

func (q *Queries) SetUsersPass(ctx context.Context, arg SetUsersPassParams) error {
	_, err := q.db.ExecContext(ctx, setUsersPass, arg.ID, arg.Password)
	return err
}