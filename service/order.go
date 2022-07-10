package service

import (
	"github.com/hsuehly/presideService/db"
	"github.com/hsuehly/presideService/models"
)

func CreatOrderService(username, phone string, identity int, weddingname, city, address, remarks string) bool {
	cli := db.Get()
	order := models.OrderDetails{
		Username:    username,
		Phone:       phone,
		Identity:    identity,
		Weddingname: weddingname,
		City:        city,
		Address:     address,
		Remarks:     remarks,
	}
	tx := cli.Create(&order)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
func GetOrderService() (orderDetails []*models.OrderDetails) {
	db.Get().Find(&orderDetails)
	return orderDetails

}

func GetOrderByIdService(id int) (orderDetails []*models.OrderDetails) {
	db.Get().First(&orderDetails, id)
	return
}
func UpdateOrderByIdService(id int, username, phone string) (orderDetail models.OrderDetails, ok bool, err error) {
	err = db.Get().First(&orderDetail, id).Error
	if err != nil {
		return orderDetail, false, err
	}
	orderDetail.Username = username
	orderDetail.Phone = phone
	err = db.Get().Save(&orderDetail).Error
	if err != nil {
		return orderDetail, false, err
	}
	return orderDetail, true, nil
}
