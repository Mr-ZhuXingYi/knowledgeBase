package dto

import "time"

type OrderList struct {
	Id          int       `json:"id" `
	OrderNum    string    `json:"order_num"`
	CreateTime  time.Time `json:"create_time"`
	Status      string    `json:"status"`
	OrderAmount float64   `json:"order_amount"`
	CouponCode  string    `json:"coupon_code"`
	UserId      int       `json:"user_id" json:"user_id" `
	UpdateTime  time.Time `json:"update_time"`
	GoodsId     int       `json:"goods_id"`
	GoodsName   int       `json:"goods_name"`
	GoodsPrices float64   `json:"goods_prices"`
	GoodsCount  int       `json:"goods_count"`
}

type OrderReq struct {
	OrderAmount float64 `json:"order_amount"`
	CouponCode  string  `json:"coupon_code"`
	UserId      int     `json:"user_id" `
	GoodsId     int     `json:"goods_id"`
	GoodsName   string  `json:"goods_name"`
	GoodsPrices float64 `json:"goods_prices"`
	GoodsCount  int     `json:"goods_count"`
}
