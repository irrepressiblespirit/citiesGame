package game_algorithm

import (
	"github.com/skibnev/citiesGame/models"
	"strings"
)

func getLastChar(city string) rune {
	cityWithoutHyphen := strings.Trim(city, "-")
	cityWithoutSpaces := strings.Replace(cityWithoutHyphen, " ", "", -1)
	ind := (len(cityWithoutSpaces) / 2) - 1
	lastChar := []rune(cityWithoutSpaces)[ind]
	switch lastChar {
	case 'й':
		return 'и'
	case 'ь':
		return []rune(cityWithoutSpaces)[(len(cityWithoutSpaces)/2)-2]
	case 'ы':
		return []rune(cityWithoutSpaces)[(len(cityWithoutSpaces)/2)-2]
	case 'ё':
		return 'е'
	case 'щ':
		return 'ш'
	default:
		return lastChar
	}
}

func getFirstChar(city string) rune {
	return []rune(city)[0]
}

func GetNewCity(userCity string, cities map[string]string, user *models.UserInfo) bool {
	firstChar := getLastChar(userCity)
	for city := range cities {
		if getFirstChar(city) == firstChar {
			user.UsedCities[strings.ToLower(userCity)] = nil
			user.UsedCities[strings.ToLower(city)] = nil
			user.PrevGameAnswer = &models.GameAnswer{
				LastChar: getLastChar(city),
				NewCity:  city,
			}
			return true
		}
	}
	return false
}
