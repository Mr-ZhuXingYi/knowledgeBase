package configuration

import (
	"jtthink.base/src/dao"
	"jtthink.base/src/service/OrderService"
)

type ServiceConfig struct{}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (this *ServiceConfig) OrderDao() *dao.OrderDao {
	return dao.NewOrderDao()
}

func (this *ServiceConfig) OrderService() *OrderService.OrderService {
	return OrderService.NewOrderService()
}

func (this *ServiceConfig) SubOrderDao() *dao.SubOrderDao {
	return dao.NewSubOrderDao()
}
