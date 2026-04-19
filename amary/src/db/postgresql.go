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

type DBTX interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type DBHandle struct {
	db     *sql.DB
	logger Logger
	isProd bool
}

func (cdb *DBHandle) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	if !cdb.isProd {
		cdb.logger.Infoln(query)
	}
	return cdb.db.QueryContext(ctx, query, args...)
}

func (cdb *DBHandle) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	if !cdb.isProd {
		cdb.logger.Infoln(query)
	}
	return cdb.db.QueryRowContext(ctx, query, args...)
}

func (cdb *DBHandle) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	if !cdb.isProd {
		cdb.logger.Infoln(query)
	}
	return cdb.db.ExecContext(ctx, query, args...)
}

func (cdb *DBHandle) Close() error {
	return cdb.db.Close()
}

func (cdb *DBHandle) Begin(ctx context.Context, opt *sql.TxOptions) (*sql.Tx, error) {
	return cdb.db.BeginTx(ctx, opt)
}

func wrapDB(db *sql.DB, logger Logger, isProd bool) *DBHandle {
	return &DBHandle{db, logger, isProd}
}

func ConnectDB(logger Logger, isProd bool) (*DBHandle, error) {
	dbHost := os.Getenv("AMARY_DATABASE_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("AMARY_POSTGRES_DB")
	dbUser := os.Getenv("AMARY_POSTGRES_USER")
	dbPass := os.Getenv("AMARY_POSTGRES_PASSWORD")

	connString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	conn, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, err
	}
	return wrapDB(conn, logger, isProd), nil
}
