package helper

import (
	"github.com/skibnev/citiesGame/models"
)

var data = make(map[int64]models.UserInfo)

func GetData(chatId int64) models.UserInfo {
	newUser := data[chatId]
	return newUser

}

func PutData(chatId int64, userData models.UserInfo) {
	data[chatId] = userData
}
