package configuration

import (
	"jtthink.base/src/dao"
	"jtthink.base/src/service"
)

type ServiceConfig struct{}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (this *ServiceConfig) OrderDao() *dao.OrderDao {
	return dao.NewOrderDao()
}

func (this *ServiceConfig) OrderService() *service.OrderService {
	return service.NewOrderService()
}

func (this *ServiceConfig) SubOrderDao() *dao.SubOrderDao {
	return dao.NewSubOrderDao()
}
