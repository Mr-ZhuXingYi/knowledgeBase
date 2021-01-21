package assembler

import (
	"github.com/shopspring/decimal"
	"jtthink.base/src/application/dto"
	"jtthink.base/src/models/OrderModel"
	"jtthink.base/src/models/SubOrderModel"
	"jtthink.base/src/service"
	"jtthink.base/util"
	"time"
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

func D2M_Order(req *dto.OrderReq) *OrderModel.Orders {
	return &OrderModel.Orders{
		OrderNum:    util.GetOrderNum(),
		UserId:      req.UserId,
		CreateTime:  time.Now(),
		Status:      service.STATUS_HAVE_PAY,
		OrderAmount: decimal.NewFromFloat(req.OrderAmount),
		CouponCode:  req.CouponCode,
		UpdateTime:  time.Now(),
	}
}

func D2M_SubOrder(req *dto.OrderReq, orderId int) *SubOrderModel.SubOrders {
	return &SubOrderModel.SubOrders{
		OrderId:     orderId,
		GoodsId:     req.GoodsId,
		GoodsName:   req.GoodsName,
		GoodsPrices: decimal.NewFromFloat(req.GoodsPrices),
		GoodsCount:  req.GoodsCount,
	}
}
