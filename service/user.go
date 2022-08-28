package service

import (
	"github.com/hsuehly/presideService/db"
	"github.com/hsuehly/presideService/models"
)

type UserApi struct {
	Username string `gorm:"column:user_name;NOT NULL;type:varchar(20)" form:"username" json:"username"`
	Images   string `gorm:"column:images;NOT NULL;type:json" form:"images" json:"images"`
	Videos   string `gorm:"column:videos;NOT NULL;type:json" form:"videos" json:"videos"`
}

func GetUserByIdService(id string) (userInfo *UserApi) {
	db.Get().Model(&models.UserInfo{}).Find(&userInfo, "user_id = ?", id)
	return
}
