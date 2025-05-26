package main

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

const userTable = `
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Users table: Stores all platform users, including company admins, employees, and accounting staff.
CREATE TABLE users (
    uid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    password TEXT NOT NULL,
    email VARCHAR(255) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
)`
const tokenTable = `
CREATE TABLE tokens (
    uid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
    updated_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
    user_uid UUID NOT NULL REFERENCES users(uid),
    token text NOT NULL UNIQUE,
    expiry timestamp with time zone NOT NULL,
    mobile boolean NOT NULL DEFAULT FALSE,
    identifier text
)`

func init() {
	up := []string{
		userTable,
		tokenTable,
	}
	down := []string{
		`DROP TABLE tokens`,
		`DROP TABLE users`,
	}
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Println("create user table")
		for _, q := range up {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Println("drop user table")
		for _, q := range down {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
