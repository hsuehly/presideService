package service

import (
	"github.com/hsuehly/presideService/db"
	"github.com/hsuehly/presideService/models"
)

func GetFoodByIdService(id string) string {
	var result string
	db.Get().Model(&models.FoodMenu{}).Select("data").First(&result, "id = ?", id)
	return result

}
