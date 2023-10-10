package types

import "database/sql"

type T_category struct {
	Category_id    int            `db:"category_id"`
	Category_name  string         `db:"category_name"`
	Parent_id      sql.NullString `db:"parent_id"`
	Category_level int            `db:"category_level"`
	Category_order int            `db:"category_order"`
}
