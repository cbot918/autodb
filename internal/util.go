package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

func PrintJSON(v any) {
	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("json marshalindent failed")
	}
	fmt.Println(string(json))
}

func NewDB(cfg *Config) (*sql.DB, error) {
	//"root:12345@tcp(localhost:3309)/sec-kill?charset=utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		cfg.DB_USER,
		cfg.DB_PASSWORD,
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_NAME,
	)

	conn, err := sql.Open(cfg.DB_DRIVER, dsn)
	if err != nil {
		return nil, err
	}
	return conn, err
}
