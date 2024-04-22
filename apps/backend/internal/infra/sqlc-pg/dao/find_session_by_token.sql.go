// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: find_session_by_token.sql

package dao

import (
	"context"
)

const findSessionByToken = `-- name: FindSessionByToken :one
SELECT sessions.token, sessions.player_id FROM "sessions" WHERE "token" = $1
`

type FindSessionByTokenRow struct {
	Session Session
}

func (q *Queries) FindSessionByToken(ctx context.Context, token string) (*FindSessionByTokenRow, error) {
	row := q.db.QueryRow(ctx, findSessionByToken, token)
	var i FindSessionByTokenRow
	err := row.Scan(&i.Session.Token, &i.Session.PlayerID)
	return &i, err
}
