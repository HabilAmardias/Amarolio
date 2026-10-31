package db

import (
	"context"
	"database/sql"
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

func (cdb *DBHandle) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	if !cdb.isProd {
		cdb.logger.Infoln(query)
	}
	return cdb.db.PrepareContext(ctx, query)
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
	db, err := sql.Open("sqlite3", "./amarfin.db")
	if err != nil {
		return nil, err
	}
	return wrapDB(db, logger, isProd), nil
}
