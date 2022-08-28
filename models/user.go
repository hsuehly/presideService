package models

import (
	"github.com/hsuehly/presideService/db"
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	ID        uint           `gorm:"primarykey"`
	UserId    string         `gorm:"column:user_id;NOT NULL;type:varchar(200)" json:"userid"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Username  string         `gorm:"column:user_name;NOT NULL;type:varchar(20)" form:"username" json:"username"`
	Images    string         `gorm:"column:images;NOT NULL;type:json" form:"images" json:"images"`
	Videos    string         `gorm:"column:videos;NOT NULL;type:json" form:"videos" json:"videos"`
}

func (od *OrderDetails) UserInfo() string {
	return "user_info"

}

func init() {
	err := db.Get().AutoMigrate(&UserInfo{})
	if err != nil {
		panic("init OrderDetails failed" + err.Error())
	}

}
