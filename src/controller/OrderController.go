package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type Ordercontroller struct {
}

func NewOrdercontroller() *Ordercontroller {
	return &Ordercontroller{}
}

func (this *Ordercontroller) Name() string {
	return "OrderController"
}

func (this *Ordercontroller) Create(ctx *gin.Context) string {
	return "success"
}

func (this *Ordercontroller) Build(goft *goft.Goft) {
	goft.Handle("GET", "/order", this.Create)
}
