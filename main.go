package main

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/skibnev/citiesGame/config"
	"github.com/skibnev/citiesGame/game_algorithm"
	"github.com/skibnev/citiesGame/helper"
	"github.com/skibnev/citiesGame/models"
	"github.com/skibnev/citiesGame/validator"
	"log"
	"os"
	"reflect"
)

func main() {
	var message tgbotapi.MessageConfig
	var citiesMap map[string]string
	//uncomment when you need to add cities to the json file from the geoNames API
	//cities := cities_api.ParseCities()
	//citiesInJson, _ := json.Marshal(cities)
	//helper.WriteToFile("Cities.json", citiesInJson)
	bot, err := tgbotapi.NewBotAPI(config.TelegramKey)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			if update.Message.Text == "/start" && helper.GetData(update.Message.Chat.ID).UsedCities == nil {
				json.Unmarshal(helper.ReadFromFile("Cities.json"), &citiesMap)
				newUserData := models.UserInfo{
					LastChar:   make([]rune, 1),
					NewCity:    make([]string, 1),
					UsedCities: make(map[string]string),
				}
				helper.PutData(update.Message.Chat.ID, newUserData)
				message = tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать в игру ГОРОДА. Введите название города")
				bot.Send(message)
			} else {
				userData := helper.GetData(update.Message.Chat.ID)
				if validator.CheckUserInput(update.Message.Text, citiesMap, userData.UsedCities, &userData.LastChar[0]) {
					cityFound := game_algorithm.GetNewCity(update.Message.Text, citiesMap, &userData)
					if !cityFound {
						message = tgbotapi.NewMessage(update.Message.Chat.ID, "Поздравляем !!! Вы выиграли")
						os.Exit(0)
					} else {
						message = tgbotapi.NewMessage(update.Message.Chat.ID, userData.NewCity[0])
					}
					bot.Send(message)
				} else {
					message = tgbotapi.NewMessage(update.Message.Chat.ID, "Не корректно введенное слово !!! Попробуйте еще раз")
					bot.Send(message)
				}
			}
		}
	}
}
