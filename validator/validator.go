package validator

import (
	"github.com/skibnev/citiesGame/config"
	"github.com/skibnev/citiesGame/external_api"
	"regexp"
)

func CheckUserInput(city string, cities map[string]string, usedCityMap map[string]string, lastChar *rune) bool {
	if *lastChar == 0 {
		return checkCityNameSize(city) && !isContainsProhibitedCharacters(city) && checkCityIsExist(city, cities)
	}
	return checkCityNameSize(city) && !isContainsProhibitedCharacters(city) && isCityNameStartWithLastCharFromPrevWord(city, string(*lastChar)) && !isCityUsed(city, usedCityMap) && checkCityIsExist(city, cities)
}

func checkCityNameSize(cityName string) bool {
	return len(cityName) != 0 && len(cityName) > 0
}

func checkCityIsExist(cityName string, cities map[string]string) bool {
	cityInMap := cities[cityName]
	if len(cityInMap) != 0 {
		return true
	}
	return external_api.CheckCityIsExistByExternalApi(cityName)
}

func isCityNameStartWithLastCharFromPrevWord(city string, char string) bool {
	return string([]rune(city)[0]) == char
}

func isCityUsed(city string, usedCitiesMap map[string]string) bool {
	return len(usedCitiesMap[city]) > 0
}

func isContainsProhibitedCharacters(userCity string) bool {
	pattern, _ := regexp.Compile(config.ProhibitedCharacters)
	return pattern.Match([]byte(userCity))
}
