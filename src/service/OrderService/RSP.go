package OrderService

import (
	"jtthink.base/src/models/OrderModel"
	"jtthink.base/src/models/SubOrderModel"
	"time"
)

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

func M2D(orders []*OrderModel.Orders, subOrders []*SubOrderModel.SubOrders) []*OrderList {
	ret := make([]*OrderList, 0)
	for _, order := range orders {
		for _, subOrder := range subOrders {
			if subOrder.OrderId == order.Id {
				goodsPrices, _ := subOrder.GoodsPrices.Float64()
				orderAmount, _ := order.OrderAmount.Float64()
				status := GetStatus(order.Status)
				ret = append(ret, &OrderList{
					Id:          order.Id,
					OrderNum:    order.OrderNum,
					CreateTime:  order.CreateTime,
					Status:      status,
					OrderAmount: orderAmount,
					CouponCode:  order.CouponCode,
					UserId:      order.UserId,
					UpdateTime:  order.UpdateTime,
					GoodsId:     subOrder.GoodsId,
					GoodsName:   subOrder.GoodsName,
					GoodsPrices: goodsPrices,
					GoodsCount:  subOrder.GoodsCount,
				})
			}
		}

	}
	return ret
}

func GetStatus(i int) string {
	switch i {
	case 1:
		return "已支付"
	case 2:
		return "未支付"
	case 3:
		return "已取消"
	default:
		return ""
	}
}
