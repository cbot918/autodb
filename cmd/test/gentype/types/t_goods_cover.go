package types

type T_goods_cover struct {
	Gc_id        int    `db:"gc_id"`
	Goods_id     int    `db:"goods_id"`
	Gc_pic_url   string `db:"gc_pic_url"`
	Gc_thumb_url string `db:"gc_thumb_url"`
	Gc_order     int    `db:"gc_order"`
}
