package models

import (
	"github.com/hsuehly/presideService/db"
)

type FoodMenu struct {
	ID   uint                     `gorm:"primarykey"`
	Data []map[string]interface{} `gorm:"column:data;NOT NULL;type:json"  json:"data"`
}

func (od *OrderDetails) Food() string {
	return "food_menu"

}

func init() {
	err := db.Get().AutoMigrate(&FoodMenu{})
	if err != nil {
		panic("init UserInfo failed" + err.Error())
	}

}
