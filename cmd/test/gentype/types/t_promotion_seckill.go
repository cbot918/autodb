package types

import "database/sql"

type T_promotion_seckill struct {
	Ps_id         int64          `db:"ps_id"`
	Goods_id      int            `db:"goods_id"`
	Ps_count      int            `db:"ps_count"`
	Start_time    sql.NullString `db:"start_time"`
	End_time      sql.NullString `db:"end_time"`
	Status        sql.NullString `db:"status"`
	Current_price float32        `db:"current_price"`
	Version       int            `db:"version"`
}
