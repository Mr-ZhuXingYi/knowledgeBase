package assembler

import (
	"github.com/shopspring/decimal"
	"jtthink.base/src/application/dto"
	"jtthink.base/src/models/OrderModel"
	"jtthink.base/src/models/SubOrderModel"
	"jtthink.base/util"
	"time"
)

const (
	STATUS_ZERO           = iota
	STATUS_HAVE_PAY       //未支付
	STATUS_HAVENO_PAY     //已支付
	STATUS_HAVE_CANCELLES //已取消
)

const (
	GOODS_ZERO   = iota
	GOODS_MOUNTH //月会员
	GOODS_SEASON //季度会员
	GOODS_YEARAR //年会员
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
	case STATUS_HAVE_PAY:
		return "已支付"
	case STATUS_HAVENO_PAY:
		return "未支付"
	case STATUS_HAVE_CANCELLES:
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
		Status:      STATUS_HAVE_PAY,
		OrderAmount: decimal.NewFromFloat(req.OrderAmount),
		CouponCode:  req.CouponCode,
		UpdateTime:  time.Now(),
	}
}

func D2M_SubOrder(req *dto.OrderReq, orderId int) *SubOrderModel.SubOrders {
	return &SubOrderModel.SubOrders{
		OrderId:     orderId,
		GoodsId:     req.GoodsId,
		GoodsName:   GetGoodsName(req.GoodsName),
		GoodsPrices: decimal.NewFromFloat(req.GoodsPrices),
		GoodsCount:  req.GoodsCount,
	}
}

func GetGoodsName(goods string) int {
	switch goods {
	case "月会员":
		return GOODS_MOUNTH
	case "季会员":
		return GOODS_SEASON
	case "年会员":
		return GOODS_YEARAR
	default:
		return 0
	}
}
