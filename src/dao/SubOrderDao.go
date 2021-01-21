package dao

import (
	"github.com/shenyisyn/goft-gin/goft"
	"gorm.io/gorm"
	"jtthink.base/src/models/SubOrderModel"
)

type SubOrderDao struct{}

func NewSubOrderDao() *SubOrderDao {
	return &SubOrderDao{}
}

func (this *SubOrderDao) GetSubOrderByOrderIds(db *gorm.DB, orderIds []int) []*SubOrderModel.SubOrders {
	ret := make([]*SubOrderModel.SubOrders, 0)
	goft.Error(db.Table(SubOrderModel.SubOrdersTableName).
		Where("order_id in ?", orderIds).Scan(&ret).Error)
	return ret
}

func (this *SubOrderDao) Create(db *gorm.DB, subOrder *SubOrderModel.SubOrders) {
	goft.Error(db.Table(SubOrderModel.SubOrdersTableName).Create(subOrder).Error)
}
