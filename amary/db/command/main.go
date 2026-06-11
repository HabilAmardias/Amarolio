package dbcommand

import (
	"amary/src/db"
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

type Logger interface {
	Infoln(args ...interface{})
}

type IDEncoderItf interface {
	Encode(id int64) string
}

func getFiles(dirPath string) ([]string, error) {
	filenames := []string{}

	if err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			filenames = append(filenames, path)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return filenames, nil
}

func Backfill(db *db.DBHandle, ide IDEncoderItf) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	query := `
		SELECT id 
		FROM urls 
		WHERE short_code IS NULL OR short_code = ''
		ORDER BY id ASC
	`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to fetch un-backfilled rows: %w", err)
	}
	defer rows.Close()

	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return fmt.Errorf("failed to scan row ID: %w", err)
		}
		ids = append(ids, id)
	}

	if len(ids) == 0 {
		return nil
	}

	updateStmt, err := db.PrepareContext(ctx, "UPDATE urls SET short_code = $1 WHERE id = $2")
	if err != nil {
		return fmt.Errorf("failed to prepare update statement: %w", err)
	}
	defer updateStmt.Close()

	for _, id := range ids {
		shortCode := ide.Encode(id)

		_, err := updateStmt.ExecContext(ctx, shortCode, id)
		if err != nil {
			return fmt.Errorf("failed to backfill ID %d: %w", id, err)
		}
	}
	constraintQuery := `
		CREATE UNIQUE INDEX IF NOT EXISTS unique_short_code 
        ON urls (short_code);
	`
	if _, err := db.ExecContext(ctx, constraintQuery); err != nil {
		return fmt.Errorf("failed to enforce unique constraint: %w", err)
	}
	return nil
}

func Migration(db *db.DBHandle, lg Logger) error {
	fnames, err := getFiles("./db/migration")
	if err != nil {
		return err
	}
	query := ""
	for _, fn := range fnames {
		content, err := os.ReadFile(fn)
		if err != nil {
			return err
		}
		query += string(content)
	}

	_, err = db.ExecContext(context.Background(), query)
	if err != nil {
		return err
	}
	lg.Infoln("Migration Success")
	return nil
}
