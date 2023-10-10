package types

import "time"

type T_success_killed struct {
	Goods_id    int64     `db:"goods_id"`
	User_id     int64     `db:"user_id"`
	State       int8      `db:"state"`
	Create_time time.Time `db:"create_time"`
}
