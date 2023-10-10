package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
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
	var dsn string
	if cfg.DB_DRIVER == "mysql" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			cfg.DB_USER,
			cfg.DB_PASSWORD,
			cfg.DB_HOST,
			cfg.DB_PORT,
			cfg.DB_NAME,
		)
	} else if cfg.DB_DRIVER == "postgres" {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.DB_USER,
			cfg.DB_PASSWORD,
			cfg.DB_HOST,
			cfg.DB_PORT,
			cfg.DB_NAME,
		)
	} else {
		return nil, fmt.Errorf("database driver not found, check .env")
	}

	conn, err := sql.Open(cfg.DB_DRIVER, dsn)
	if err != nil {
		return nil, err
	}
	return conn, err
}

func Init() (*Config, *sql.DB, error, error) {

	cfg, cfgErr := NewConfig()
	db, dbErr := NewDB(cfg)

	return cfg, db, cfgErr, dbErr
}

func RowsBind(rows *sql.Rows, dest interface{}) error {
	// Use reflection to get the type of the destination struct
	destType := reflect.TypeOf(dest).Elem()

	// Use reflection to get the values of the destination struct's fields
	destValues := make([]reflect.Value, destType.NumField())

	for i := 0; i < destType.NumField(); i++ {
		field := destType.Field(i)
		destValues[i] = reflect.ValueOf(dest).Elem().FieldByName(field.Name)
	}

	// Create an interface slice for Scan
	destInterfaces := make([]interface{}, len(destValues))
	for i, v := range destValues {
		destInterfaces[i] = v.Addr().Interface()
	}

	// Scan the row into the struct fields
	if err := rows.Scan(destInterfaces...); err != nil {
		return err
	}

	return nil
}

func RowBind(row *sql.Row, dest interface{}) error {
	destType := reflect.TypeOf(dest).Elem()

	// Use reflection to get the values of the destination struct's fields
	destValues := make([]reflect.Value, destType.NumField())

	for i := 0; i < destType.NumField(); i++ {
		field := destType.Field(i)
		destValues[i] = reflect.ValueOf(dest).Elem().FieldByName(field.Name)
	}

	// Create an interface slice for Scan
	destInterfaces := make([]interface{}, len(destValues))
	for i, v := range destValues {
		destInterfaces[i] = v.Addr().Interface()
	}

	// Scan the row into the struct fields
	if err := row.Scan(destInterfaces...); err != nil {
		return err
	}

	return nil
}
