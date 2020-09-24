package goose

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/pkg/errors"
)

type tmplVars struct {
	Version   string
	CamelName string
}

// CreateWithTemplate writes a new blank migration file.
func CreateWithTemplate(cfg *config, tmpl *template.Template, name, migrationType string) error {
	version := time.Now().Format(timestampFormat)
	filename := fmt.Sprintf("%v_%v.%v", version, snakeCase(name), migrationType)

	if tmpl == nil {
		if migrationType == "go" {
			tmpl = goSQLMigrationTemplate
		} else {
			tmpl = sqlMigrationTemplate
		}
	}

	path := filepath.Join(cfg.dir, filename)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return errors.Wrap(err, "failed to create migration file")
	}

	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "failed to create migration file")
	}
	defer f.Close()

	vars := tmplVars{
		Version:   version,
		CamelName: camelCase(name),
	}
	if err := tmpl.Execute(f, vars); err != nil {
		return errors.Wrap(err, "failed to execute tmpl")
	}

	log.Printf("Created new file: %s\n", f.Name())
	return nil
}

// Create writes a new blank migration file.
func Create(cfg *config, name, migrationType string) error {
	return CreateWithTemplate(cfg, nil, name, migrationType)
}

var sqlMigrationTemplate = template.Must(template.New("goose.sql-migration").Parse(`-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
`))

var goSQLMigrationTemplate = template.Must(template.New("goose.go-migration").Parse(`package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(up{{.CamelName}}, down{{.CamelName}})
}

func up{{.CamelName}}(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func down{{.CamelName}}(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
`))
