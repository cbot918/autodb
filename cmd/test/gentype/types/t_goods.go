package types

import "database/sql"

// type T_goods struct {
// 	Goods_id         int            `db:"goods_id"`
// 	Title            string         `db:"title"`
// 	Sub_title        sql.NullString `db:"sub_title"`
// 	Original_cost    float32        `db:"original_cost"`
// 	Current_price    float32        `db:"current_price"`
// 	Discount         float32        `db:"discount"`
// 	Is_free_delivery int            `db:"is_free_delivery"`
// 	Category_id      int            `db:"category_id"`
// 	Last_update_time sql.NullString `db:"last_update_time"`
// }

type T_goods struct {
	Goods_id         int
	Title            string
	Sub_title        sql.NullString
	Original_cost    float32
	Current_price    float32
	Discount         float32
	Is_free_delivery int
	Category_id      int
	Last_update_time sql.NullString
}
