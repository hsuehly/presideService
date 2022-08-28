package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hsuehly/presideService/Response"
	"github.com/hsuehly/presideService/models"
	"github.com/hsuehly/presideService/service"
	"github.com/hsuehly/presideService/util"
	"strconv"
)

func CreateOrder(c *gin.Context) {
	var orderDetail models.OrderDetails
	//orderDetail := make(map[string]interface{}, 7)
	userId := c.Param("id")
	err := c.ShouldBind(&orderDetail)
	if err != nil {
		Response.ValidatorErr(c, err)
		return
	}
	var TempD = make(map[string]*util.TempData)
	TempD["username"] = &util.TempData{
		Value: orderDetail.Username,
		Color: "#173177",
	}
	TempD["phone"] = &util.TempData{
		Value: orderDetail.Phone,
		Color: "#173177",
	}
	TempD["identity"] = &util.TempData{
		Value: orderDetail.Identity,
		Color: "#173177",
	}
	TempD["weddingname"] = &util.TempData{
		Value: orderDetail.Weddingname,
		Color: "#173177",
	}
	TempD["date"] = &util.TempData{
		Value: orderDetail.Date,
		Color: "#173177",
	}
	TempD["times"] = &util.TempData{
		Value: orderDetail.Times,
		Color: "#173177",
	}
	TempD["city"] = &util.TempData{
		Value: orderDetail.City,
		Color: "#173177",
	}
	TempD["address"] = &util.TempData{
		Value: orderDetail.Address,
		Color: "#173177",
	}
	TempD["remarks"] = &util.TempData{
		Value: orderDetail.Remarks,
		Color: "#173177",
	}
	go util.SendTemp(TempD)
	ok := service.CreatOrderService(userId, orderDetail.Username, orderDetail.Phone, orderDetail.Identity, orderDetail.Weddingname, orderDetail.City, orderDetail.Address, orderDetail.Remarks, orderDetail.Times, orderDetail.Date)
	if !ok {
		Response.Error(c, Response.ApiCode.CREATEORDERFAILED, Response.ApiCode.GetMessage(Response.ApiCode.CREATEORDERFAILED))
		return
	}
	Response.Success(c, userId)

}
func GetOrder(c *gin.Context) {
	orderList := service.GetOrderService()
	fmt.Println("orderList", orderList)
	Response.Success(c, orderList)
}

func GetOrderById(c *gin.Context) {
	id := c.Param("id")
	//if !ok {
	//	util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
	//	return
	//}

	order := service.GetOrderByIdService(id)

	Response.Success(c, order)
}
func UpdateOrderById(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		Response.Error(c, Response.ApiCode.LCAKPARAMETERS, Response.ApiCode.GetMessage(Response.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id)
	if err != nil {
		Response.Error(c, Response.ApiCode.CONVERTFAILED, Response.ApiCode.GetMessage(Response.ApiCode.CONVERTFAILED))
		return
	}

	username, ok := c.GetPostForm("username")
	if !ok {
		Response.Error(c, Response.ApiCode.LCAKPARAMETERS, Response.ApiCode.GetMessage(Response.ApiCode.LCAKPARAMETERS))
		return
	}

	phone, ok := c.GetPostForm("phone")
	if !ok {
		Response.Error(c, Response.ApiCode.LCAKPARAMETERS, Response.ApiCode.GetMessage(Response.ApiCode.LCAKPARAMETERS))
		return
	}
	if err != nil {
		Response.Error(c, Response.ApiCode.CONVERTFAILED, Response.ApiCode.GetMessage(Response.ApiCode.CONVERTFAILED))
		return
	}

	_, ok, err = service.UpdateOrderByIdService(newId, username, phone)
	if !ok {
		Response.Error(c, Response.ApiCode.NOSUCHID, Response.ApiCode.GetMessage(Response.ApiCode.NOSUCHID))
		return
	}

	if err != nil {
		Response.Error(c, Response.ApiCode.FAILED, Response.ApiCode.GetMessage(Response.ApiCode.FAILED))
		return
	}

	Response.Success(c, nil)
}
