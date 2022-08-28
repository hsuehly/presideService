package models

import (
	"github.com/hsuehly/presideService/db"
	"gorm.io/gorm"
	"time"
)

type OrderDetails struct {
	ID          uint           `gorm:"primarykey"`
	UserId      string         `gorm:"column:user_id;NOT NULL;type:varchar(200)" json:"userid"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Username    string         `gorm:"column:user_name;NOT NULL;type:varchar(20)" form:"username" json:"username" binding:"required"`
	Phone       string         `gorm:"column:user_phone;NOT NULL;type:varchar(100)" form:"phone" json:"phone"  binding:"required,min=11"`
	Identity    string         `gorm:"column:identity;NOT NULL;type:varchar(100)" form:"identity" json:"identity" binding:"required"`
	Weddingname string         `gorm:"column:weddingname;type:varchar(100)" form:"weddingname"  json:"weddingname"`
	Date        string         `gorm:"column:date;NOT NULL;type:varchar(100)" form:"date"  json:"date" binding:"required"`
	Times       string         `gorm:"column:times;NOT NULL;type:varchar(10)" form:"times" json:"times" binding:"required"`
	City        string         `gorm:"column:city;NOT NULL;type:varchar(255)" form:"city"  json:"city" binding:"required"`
	Address     string         `gorm:"column:address;type:varchar(200)" form:"address"  json:"address" `
	Remarks     string         `gorm:"column:remarks;type:varchar(100)" form:"remarks"  json:"remarks"`
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
