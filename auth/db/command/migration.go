package dbcommand

import (
	"amarolio-auth/src/db"
	"context"
	"io/fs"
	"os"
	"path/filepath"
)

type Logger interface {
	Infoln(args ...interface{})
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
