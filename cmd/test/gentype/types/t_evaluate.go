package types

import "database/sql"

type T_evaluate struct {
	Evaluate_id int64          `db:"evaluate_id"`
	Content     string         `db:"content"`
	Stars       sql.NullString `db:"stars"`
	Create_time sql.NullString `db:"create_time"`
	Goods_id    int64          `db:"goods_id"`
}
