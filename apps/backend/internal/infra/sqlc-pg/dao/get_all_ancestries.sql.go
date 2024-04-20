// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: get_all_ancestries.sql

package dao

import (
	"context"
)

const getAllAncestries = `-- name: GetAllAncestries :many
SELECT ancestries.id, ancestries.code, ancestries.title FROM "ancestries"
`

type GetAllAncestriesRow struct {
	Ancestry Ancestry
}

func (q *Queries) GetAllAncestries(ctx context.Context) ([]*GetAllAncestriesRow, error) {
	rows, err := q.db.Query(ctx, getAllAncestries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetAllAncestriesRow
	for rows.Next() {
		var i GetAllAncestriesRow
		if err := rows.Scan(&i.Ancestry.ID, &i.Ancestry.Code, &i.Ancestry.Title); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}