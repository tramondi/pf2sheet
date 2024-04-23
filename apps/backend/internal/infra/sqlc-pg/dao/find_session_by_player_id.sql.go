// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: find_session_by_player_id.sql

package dao

import (
	"context"
)

const findSessionByPlayerID = `-- name: FindSessionByPlayerID :one
SELECT sessions.token, sessions.player_id FROM "sessions" WHERE "player_id" = $1
`

type FindSessionByPlayerIDRow struct {
	Session Session
}

func (q *Queries) FindSessionByPlayerID(ctx context.Context, playerID int32) (*FindSessionByPlayerIDRow, error) {
	row := q.db.QueryRow(ctx, findSessionByPlayerID, playerID)
	var i FindSessionByPlayerIDRow
	err := row.Scan(&i.Session.Token, &i.Session.PlayerID)
	return &i, err
}
