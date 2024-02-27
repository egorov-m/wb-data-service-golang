package database

import "github.com/jackc/pgx/v5"

func CollectOneRow[T any](rows pgx.Rows) (*T, error) {
	return pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByNameLax[T])
}

func CollectRows[T any](rows pgx.Rows) ([]*T, error) {
	return pgx.CollectRows(rows, pgx.RowToAddrOfStructByNameLax[T])
}
