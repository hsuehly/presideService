package service

import (
	"github.com/hsuehly/presideService/db"
	"github.com/hsuehly/presideService/models"
)

type orderApi struct {
	Username    string `gorm:"column:user_name" json:"username"`
	Phone       string `gorm:"column:user_phone" json:"phone" `
	Identity    string `gorm:"column:identity" json:"identity"`
	Weddingname string `gorm:"column:weddingname" json:"weddingname"`
	City        string `gorm:"column:city" json:"city"`
	Address     string `gorm:"column:address" json:"address"`
	Remarks     string `gorm:"column:remarks" json:"remarks"`
	Times       string `gorm:"column:times" form:"times" json:"times" `
	Date        string `gorm:"column:date" form:"date"  json:"date"`
}

func CreatOrderService(userid, username, phone string, identity, weddingname, city, address, remarks, times, data string) bool {
	cli := db.Get()
	order := models.OrderDetails{
		UserId:      userid,
		Username:    username,
		Phone:       phone,
		Identity:    identity,
		Weddingname: weddingname,
		City:        city,
		Address:     address,
		Remarks:     remarks,
		Times:       times,
		Date:        data,
	}
	tx := cli.Create(&order)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
func GetOrderService() (orderDetails []*orderApi) {
	db.Get().Model(&models.OrderDetails{}).Find(&orderDetails)
	return orderDetails

}

func GetOrderByIdService(id string) (orderDetails []*orderApi) {
	db.Get().Model(&models.OrderDetails{}).Find(&orderDetails, "user_id = ?", id)
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
