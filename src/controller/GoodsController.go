package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"gorm.io/gorm"
	"jtthink.base/src/models/GoodsModel"
)

type GoodsController struct {
	DB *gorm.DB `inject:"-"`
}

func NewGoodsController() *GoodsController {
	return &GoodsController{}
}

func (this *GoodsController) Name() string {
	return "GoodsController"
}

func (this *GoodsController) GoodsList(ctx *gin.Context) goft.Json {
	var total int64
	goodsList := []*GoodsModel.Goods{}
	t := this.DB.Table(GoodsModel.GoodsTableName)
	goft.Error(t.Scan(&goodsList).Error)
	goft.Error(t.Count(&total).Error)
	return gin.H{"goodsList": goodsList, "total": total}
}

func (this *GoodsController) Build(goft *goft.Goft) {
	goft.Handle("GET", "/goods", this.GoodsList)
}
