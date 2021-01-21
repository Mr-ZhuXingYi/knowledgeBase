package assembler

import (
	"jtthink.base/src/application/dto"
	"jtthink.base/src/models/OrderModel"
	"jtthink.base/src/models/SubOrderModel"
	"jtthink.base/src/service"
)

func M2D_OrderList(orders []*OrderModel.Orders, subOrders []*SubOrderModel.SubOrders) []*dto.OrderList {
	ret := make([]*dto.OrderList, 0)
	for _, order := range orders {
		for _, subOrder := range subOrders {
			if subOrder.OrderId == order.Id {
				goodsPrices, _ := subOrder.GoodsPrices.Float64()
				orderAmount, _ := order.OrderAmount.Float64()
				status := GetStatus(order.Status)
				ret = append(ret, &dto.OrderList{
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
	case service.STATUS_HAVE_PAY:
		return "已支付"
	case service.STATUS_HAVENO_PAY:
		return "未支付"
	case service.STATUS_HAVE_CANCELLES:
		return "已取消"
	default:
		return ""
	}
}
