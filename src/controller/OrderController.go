package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"jtthink.base/src/application/dto"
	"jtthink.base/src/service"
	"strconv"
)

const (
	GET_USERID_KEY = "userId"
)

type OrderController struct {
	OrderService *service.OrderService `inject:"-"`
}

func NewOrderController() *OrderController {
	return &OrderController{}
}

func (this *OrderController) Name() string {
	return "OrderController"
}

func (this *OrderController) Create(ctx *gin.Context) string {
	params := new(dto.OrderReq)
	goft.Error(ctx.BindJSON(params), "BindJSON err")
	userId := GetUserId(ctx)
	params.UserId = userId

	this.OrderService.Create(params)
	return "success"
}

func (this *OrderController) OrderList(ctx *gin.Context) goft.Json {
	userId := GetUserId(ctx)
	params := new(struct {
		PageNum  int `form:"page_num"`
		PageSize int `form:"page_size"`
		Key      int `form:"key"`
	})
	goft.Error(ctx.BindQuery(params), "BindQuery err")
	ret, total := this.OrderService.OrderList(params.PageNum, params.PageSize, params.Key, userId)
	return gin.H{"data": ret, "total": total}
}

func (this *OrderController) Build(goft *goft.Goft) {
	goft.Handle("GET", "/order", this.OrderList).
		Handle("POST", "/order", this.Create)
}

func GetUserId(ctx *gin.Context) int {
	userIdStr := ctx.Request.Header.Get(GET_USERID_KEY)
	userId, _ := strconv.Atoi(userIdStr)
	return userId
}
