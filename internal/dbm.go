package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBMetadata struct {
	Conn          *sql.DB
	Name          string
	TableMetadata []TableMetadata
}

type TableMetadata struct {
	Name           string
	PK             string
	ColumnMetadata []ColumnMetadata
}

type ColumnMetadata struct {
	Name     string
	DBType   string
	Nullable string
}

func NewDBMetadata(cfg *Config, conn *sql.DB) (*DBMetadata, error) {

	// initialize
	dbm := new(DBMetadata)

	dbm.Conn = conn
	dbm.Name = cfg.DB_NAME
	dbm.TableMetadata = []TableMetadata{}

	var err error

	// process
	err = dbm.GetTableName()
	if err != nil {
		return nil, err
	}

	err = dbm.GetTablePK()
	if err != nil {
		return nil, err
	}

	err = dbm.GetColumns()
	if err != nil {
		return nil, err
	}

	return dbm, nil
}

func (dbm *DBMetadata) GetTableName() error {
	q := fmt.Sprintf(`
SELECT TABLE_NAME
FROM INFORMATION_SCHEMA.TABLES
WHERE TABLE_SCHEMA = '%s';	
`, dbm.Name)
	rows, err := dbm.Conn.Query(q)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			fmt.Println("rows.Scan failed")
			return err
		}
		// fmt.Println(name)
		dbm.TableMetadata = append(dbm.TableMetadata, TableMetadata{
			Name: name,
		})
	}

	return nil
}

func (dbm *DBMetadata) GetTablePK() error {
	for index, table := range dbm.TableMetadata {
		q := fmt.Sprintf(`
SELECT COLUMN_NAME
FROM INFORMATION_SCHEMA.COLUMNS
WHERE TABLE_SCHEMA = '%s'
AND TABLE_NAME = '%s'
AND COLUMN_KEY = 'PRI'
	`, dbm.Name, table.Name)

		rows, err := dbm.Conn.Query(q)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var pk string
			if err := rows.Scan(&pk); err != nil {
				return err
			}
			dbm.TableMetadata[index].PK = pk
		}
	}
	return nil
}

func (dbm *DBMetadata) GetColumns() error {

	for index, table := range dbm.TableMetadata {
		q := fmt.Sprintf(`
SELECT COLUMN_NAME, DATA_TYPE, IS_NULLABLE
FROM INFORMATION_SCHEMA.COLUMNS
WHERE TABLE_NAME = '%s'
ORDER BY ORDINAL_POSITION;
	`, table.Name)

		rows, err := dbm.Conn.Query(q)
		if err != nil {
			return err
		}

		for rows.Next() {
			cm := ColumnMetadata{}
			err = rows.Scan(&cm.Name, &cm.DBType, &cm.Nullable)
			if err != nil {
				return err
			}
			// fmt.Println(cm)
			dbm.TableMetadata[index].ColumnMetadata = append(
				dbm.TableMetadata[index].ColumnMetadata,
				cm,
			)
		}

	}
	return nil
}
