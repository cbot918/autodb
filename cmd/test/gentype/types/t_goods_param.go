package types

type T_goods_param struct {
Gp_id int `db:"gp_id"`
	Gp_param_name string `db:"gp_param_name"`
	Gp_param_value string `db:"gp_param_value"`
	Goods_id int `db:"goods_id"`
	Gp_order int `db:"gp_order"`
	
}