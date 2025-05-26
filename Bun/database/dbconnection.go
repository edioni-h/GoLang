package database

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func dbconnection() (*bun.DB, error) {
	// Initialize the database connection here
	// For example, using PostgreSQL:
	dsn := "postgres://postgres:edioni03@localhost:5432/Studying?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	//
	if err := CheckConn(db); err != nil {
		return nil, err
	}

	return db, nil
}

func CheckConn(db *bun.DB) error {
	var n int
	return db.NewSelect().ColumnExpr("1").Scan(context.Background(), &n)
}
