package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hsuehly/presideService/service"
	"github.com/hsuehly/presideService/types"
	"github.com/hsuehly/presideService/util"
	"strconv"
)

func CreateOrder(c *gin.Context) {
	//var orderDetail models.OrderDetails
	//orderDetail := make(map[string]interface{}, 7)
	//err := c.BindJSON(&orderDetail)
	//if err != nil {
	//	util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
	//	return
	//}
	username, ok := c.GetPostForm("username")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	phone, ok := c.GetPostForm("phone")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	identity, ok := c.GetPostForm("identity")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}
	newIdentity, err := strconv.Atoi(identity)
	if err != nil {
		util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}
	weddingname, ok := c.GetPostForm("weddingname")
	if !ok && newIdentity == 1 {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}
	city, ok := c.GetPostForm("city")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}
	address := c.PostForm("address")

	remarks := c.PostForm("remarks")

	ok = service.CreatOrderService(username, phone, newIdentity, weddingname, city, address, remarks)
	if !ok {
		util.Error(c, int(types.ApiCode.CREATEORDERFAILED), types.ApiCode.GetMessage(types.ApiCode.CREATEORDERFAILED))
		return
	}
	util.Success(c, nil)

}
func GetOrder(c *gin.Context) {
	orderList := service.GetOrderService()
	fmt.Println(orderList)

	util.Success(c, orderList)
}

func GetOrderById(c *gin.Context) {
	id := c.Param("id")
	//if !ok {
	//	util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
	//	return
	//}

	newId, err := strconv.Atoi(id)
	if err != nil {
		util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	order := service.GetOrderByIdService(newId)

	util.Success(c, order)
}
func UpdateOrderById(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id)
	if err != nil {
		util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	username, ok := c.GetPostForm("username")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	phone, ok := c.GetPostForm("phone")
	if !ok {
		util.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}
	if err != nil {
		util.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	_, ok, err = service.UpdateOrderByIdService(newId, username, phone)
	if !ok {
		util.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	}

	if err != nil {
		util.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}

	util.Success(c, nil)
}
