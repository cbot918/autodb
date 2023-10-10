package types

import "database/sql"

type T_goods_detail struct {
	Gd_id      int            `db:"gd_id"`
	Goods_id   int            `db:"goods_id"`
	Gd_pic_url string         `db:"gd_pic_url"`
	Gd_order   sql.NullString `db:"gd_order"`
}
