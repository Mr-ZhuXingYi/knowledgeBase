package service

import (
	"gorm.io/gorm"
	"jtthink.base/src/application/assembler"
	"jtthink.base/src/application/dto"
	"jtthink.base/src/dao"
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
	//获取order列表
	orders, total := this.OrderDao.GetOrderByUserIdAndStatus(this.DB, pageNum, pageSize, status, userId)

	//获取subOrder列表
	orderIds := []int{}
	for _, order := range orders {
		orderIds = append(orderIds, order.Id)
	}
	subOrders := this.SubOrderDao.GetSubOrderByOrderIds(this.DB, orderIds)

	//拼装订单列表
	ret := assembler.M2D_OrderList(orders, subOrders)

	return ret, total
}

func (this *OrderService) Create(req *dto.OrderReq) {
	db := this.DB.Begin()

	//新建order
	order := assembler.D2M_Order(req)
	this.OrderDao.Create(db, order)

	//新建subOrder
	subOrder := assembler.D2M_SubOrder(req, order.Id)
	this.SubOrderDao.Create(db, subOrder)

	db.Commit()
}
