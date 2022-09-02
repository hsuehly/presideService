package service

import (
	"encoding/json"
	"fmt"
	"github.com/hsuehly/presideService/db"
	"github.com/hsuehly/presideService/models"
)

type FoodMenu struct {
	FoodTime     string `json:"foodTime"`
	FoodName     string `json:"foodName"`
	FoodImageURl string `json:"foodImageURL"`
}

func GetFoodByIdService(id string) (foodMenu []FoodMenu) {
	var result string
	db.Get().Model(&models.FoodMenu{}).Select("data").First(&result, "id = ?", id)
	err := json.Unmarshal([]byte(result), &foodMenu)
	if err != nil {
		fmt.Println("err1", err)
	}

	return

}
