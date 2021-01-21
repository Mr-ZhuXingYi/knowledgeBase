package OrderService

type OrderREQ struct {
	OrderAmount float64 `json:"order_amount"`
	CouponCode  string  `json:"coupon_code"`
	UserId      int     `json:"user_id" `
	GoodsId     int     `json:"goods_id"`
	GoodsName   int     `json:"goods_name"`
	GoodsPrices float64 `json:"goods_prices"`
	GoodsCount  int     `json:"goods_count"`
}
