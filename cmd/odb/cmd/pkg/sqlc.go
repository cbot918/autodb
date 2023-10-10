package pkg

import (
	"os"
	"os/exec"
)

func SqlcInit() error {
	// sqlc.yaml
	// -db
	var err error
	err = createFile("sqlc.yaml", sqlcConfig())
	if err != nil {
		return err
	}

	err = createFolder("db/schema")
	if err != nil {
		return err
	}

	// time.Sleep(2 * time.Second)

	err = createFile("db/schema/author.sql", schemaSample())
	if err != nil {
		return err
	}
	// time.Sleep(2 * time.Second)
	err = createFolder("db/query")
	if err != nil {
		return err
	}
	// time.Sleep(2 * time.Second)
	err = createFile("db/query/author.sql", querySample())
	if err != nil {
		return err
	}

	return exec.Command("sqlc", "generate").Run()
}

func createFolder(path string) error {

	err := os.MkdirAll(path, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func schemaSample() string {
	return `
CREATE TABLE authors (
	id   BIGSERIAL PRIMARY KEY,
	name text      NOT NULL,
	bio  text
);
`
}

func querySample() string {
	return `
-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (
	name, bio
) VALUES (
	$1, $2
)
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;
`
}

func createFile(name string, content string) error {
	fd, err := os.Create(name)
	if err != nil {
		return err
	}
	defer fd.Close()

	_, err = fd.Write([]byte(content))
	if err != nil {
		return err
	}
	return nil
}

func sqlcConfig() string {
	return `
version: "2"
sql:
  - engine: "postgresql"
    queries: "db/query/"
    schema: "db/schema/"
    # emit_json_tags: true
    # emit_prepared_queries: false
    # emit_interface: true
    # emit_exact_table_names: false
    # emit_empty_slices: true
    gen:
      go:
        package: "db"
        out: "db/sqlc"
`
}
