package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Logger interface {
	Infoln(args ...interface{})
}

type DBHandle struct {
	db     *sql.DB
	logger Logger
}

func (cdb *DBHandle) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	cdb.logger.Infoln(query)
	return cdb.db.QueryContext(ctx, query, args...)
}

func (cdb *DBHandle) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	cdb.logger.Infoln(query)
	return cdb.db.QueryRowContext(ctx, query, args...)
}

func (cdb *DBHandle) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	cdb.logger.Infoln(query)
	return cdb.db.ExecContext(ctx, query, args...)
}

func (cdb *DBHandle) Close() error {
	return cdb.db.Close()
}

func (cdb *DBHandle) Begin() (*sql.Tx, error) {
	return cdb.db.Begin()
}

func wrapDB(db *sql.DB, logger Logger) *DBHandle {
	return &DBHandle{db, logger}
}

func ConnectDB(logger Logger) (*DBHandle, error) {
	dbHost := os.Getenv("AUTH_DATABASE_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")

	connString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	conn, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, err
	}
	return wrapDB(conn, logger), nil
}
