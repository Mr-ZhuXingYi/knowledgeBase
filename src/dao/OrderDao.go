package dao

import (
	"github.com/shenyisyn/goft-gin/goft"
	"gorm.io/gorm"
	"jtthink.base/src/models/OrderModel"
)

type OrderDao struct{}

func NewOrderDao() *OrderDao {
	return &OrderDao{}
}

func (this *OrderDao) GetOrderByUserIdAndStatus(db *gorm.DB, pageNum, pageSize, status, userId int) ([]*OrderModel.Orders, int64) {
	var total int64
	ret := make([]*OrderModel.Orders, 0)
	t := db.Table(OrderModel.OrdersTableName).
		Where("user_id = ?", userId).Where("status = ?", status)

	goft.Error(t.Count(&total).Error)

	if pageNum > 0 && pageSize > 0 {
		t = t.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}

	goft.Error(t.Scan(&ret).Error)

	return ret, total
}

func (this *OrderDao) Create(db *gorm.DB, order *OrderModel.Orders) {
	goft.Error(db.Table(OrderModel.OrdersTableName).Create(order).Error)
}
