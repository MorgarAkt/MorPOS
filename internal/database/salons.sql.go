// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: salons.sql

package database

import (
	"context"
)

const createSalon = `-- name: CreateSalon :one
INSERT INTO salons (id, name)
VALUES (?, ?) RETURNING id, name
`

type CreateSalonParams struct {
	ID   string
	Name string
}

func (q *Queries) CreateSalon(ctx context.Context, arg CreateSalonParams) (Salon, error) {
	row := q.db.QueryRowContext(ctx, createSalon, arg.ID, arg.Name)
	var i Salon
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteSalon = `-- name: DeleteSalon :exec
DELETE FROM salons WHERE id = ?
`

func (q *Queries) DeleteSalon(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteSalon, id)
	return err
}

const getSalon = `-- name: GetSalon :one
SELECT id, name FROM salons WHERE id = ? LIMIT 1
`

func (q *Queries) GetSalon(ctx context.Context, id string) (Salon, error) {
	row := q.db.QueryRowContext(ctx, getSalon, id)
	var i Salon
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const listSalons = `-- name: ListSalons :many
SELECT id, name FROM salons ORDER BY id
`

func (q *Queries) ListSalons(ctx context.Context) ([]Salon, error) {
	rows, err := q.db.QueryContext(ctx, listSalons)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Salon
	for rows.Next() {
		var i Salon
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSalon = `-- name: UpdateSalon :exec
UPDATE salons
SET name = ?
WHERE id = ?
`

type UpdateSalonParams struct {
	Name string
	ID   string
}

func (q *Queries) UpdateSalon(ctx context.Context, arg UpdateSalonParams) error {
	_, err := q.db.ExecContext(ctx, updateSalon, arg.Name, arg.ID)
	return err
}