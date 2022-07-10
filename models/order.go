package models

import (
	"github.com/hsuehly/presideService/db"
)

type OrderDetails struct {
	ID          int    `gorm:"primarykey" json:"id"`
	Username    string `gorm:"column:username;NOT NULL" json:"username"`
	Phone       string `gorm:"column:phone;NOT NULL" json:"phone"`
	Identity    int    `gorm:"column:identity;NOT NULL" json:"identity"`
	Weddingname string `gorm:"column:weddingname" json:"weddingname"`
	City        string `gorm:"column:city;NOT NULL" json:"city"`
	Address     string `gorm:"column:address" json:"address"`
	Remarks     string `gorm:"column:remarks" json:"remarks"`
}

func (od *OrderDetails) OrderDetails() string {
	return "order_details"

}

func init() {
	err := db.Get().AutoMigrate(&OrderDetails{})
	if err != nil {
		panic("init OrderDetails failed" + err.Error())
	}

}
