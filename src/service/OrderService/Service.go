package OrderService

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"jtthink.base/src/dao"
	"jtthink.base/src/models/OrderModel"
	"jtthink.base/src/models/SubOrderModel"
	"time"
)

const (
	STATUS_HAVE_PAY       = iota //未支付
	STATUS_HAVENO_PAY            //已支付
	STATUS_HAVE_CANCELLES        //已取消
)

type OrderService struct {
	DB          *gorm.DB         `inject:"-"`
	OrderDao    *dao.OrderDao    `inject:"-"`
	SubOrderDao *dao.SubOrderDao `inject:"-"`
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (this *OrderService) OrderList(pageNum, pageSize, status, userId int) ([]*OrderList, int64) {
	orders, total := this.OrderDao.GetOrderByUserIdAndStatus(this.DB, pageNum, pageSize, status, userId)

	orderIds := []int{}

	for _, order := range orders {
		orderIds = append(orderIds, order.Id)
	}
	subOrders := this.SubOrderDao.GetSubOrderByOrderIds(this.DB, orderIds)

	ret := M2D(orders, subOrders)

	return ret, total
}

func (this *OrderService) Create(req *OrderREQ) {
	order := &OrderModel.Orders{
		OrderNum:    GetOrderNum(),
		UserId:      req.UserId,
		CreateTime:  time.Now(),
		Status:      STATUS_HAVE_PAY,
		OrderAmount: decimal.NewFromFloat(req.OrderAmount),
		CouponCode:  req.CouponCode,
		UpdateTime:  time.Now(),
	}

	this.OrderDao.Create(this.DB, order)

	subOrder := &SubOrderModel.SubOrders{
		OrderId:     order.Id,
		GoodsId:     req.GoodsId,
		GoodsName:   req.GoodsName,
		GoodsPrices: decimal.NewFromFloat(req.GoodsPrices),
		GoodsCount:  req.GoodsCount,
	}

	this.SubOrderDao.Create(this.DB, subOrder)
}
