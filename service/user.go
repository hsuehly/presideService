package service

import (
	"encoding/json"

	"github.com/hsuehly/presideService/db"
	"github.com/hsuehly/presideService/models"
)

type Urls struct {
	Url string `json:"url"`
}
type UserApi struct {
	Username string `gorm:"column:user_name"  json:"username"`
	Images   string `gorm:"column:images" json:"images"`
	Videos   string `gorm:"column:videos" json:"videos"`
}
type UserInfo struct {
	Username string ` json:"username"`
	Images   []Urls `json:"images"`
	Videos   []Urls `json:"videos"`
}

func GetUserByIdService(id string) (userInfo UserInfo, err error) {
	userApi := UserApi{}
	db.Get().Model(&models.UserInfo{}).Find(&userApi, "user_id = ?", id)
	var images []Urls
	var video []Urls
	err = json.Unmarshal([]byte(userApi.Images), &images)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(userApi.Videos), &video)
	if err != nil {
		return
	}

	userInfo.Username = userApi.Username
	userInfo.Images = images
	userInfo.Videos = video
	return userInfo, nil

}
