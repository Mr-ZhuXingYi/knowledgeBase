package service

import (
	"gorm.io/gorm"
	"jtthink.base/src/application/assembler"
	"jtthink.base/src/application/dto"
	"jtthink.base/src/dao"
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

func (this *OrderService) OrderList(pageNum, pageSize, status, userId int) ([]*dto.OrderList, int64) {
	orders, total := this.OrderDao.GetOrderByUserIdAndStatus(this.DB, pageNum, pageSize, status, userId)

	orderIds := []int{}

	for _, order := range orders {
		orderIds = append(orderIds, order.Id)
	}
	subOrders := this.SubOrderDao.GetSubOrderByOrderIds(this.DB, orderIds)

	ret := assembler.M2D_OrderList(orders, subOrders)

	return ret, total
}

func (this *OrderService) Create(req *dto.OrderReq) {
	order := assembler.D2M_Order(req)

	this.OrderDao.Create(this.DB, order)

	subOrder := assembler.D2M_SubOrder(req, order.Id)

	this.SubOrderDao.Create(this.DB, subOrder)
}
