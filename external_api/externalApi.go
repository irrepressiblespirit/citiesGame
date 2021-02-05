package external_api

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/skibnev/citiesGame/config"
	"golang.org/x/net/context"
	"log"
	"time"
	"googlemaps.github.io/maps"
)

type ApiResponse struct {
	TotalResultsCount float64     `json:"totalResultsCount"`
	Geonames          interface{} `json:"geonames"`
}

func CheckCityIsExistByExternalApi(cityName string) bool{
	url := fmt.Sprintf("http://api.geonames.org/searchJSON?q=%s&username=%s", cityName, config.Username)
	response, _ := resty.New().SetTimeout(3*time.Second).R().
		Get(url)
	var p *ApiResponse
	json.Unmarshal(response.Body(), &p)
	return p.TotalResultsCount > 0
}

// google service need paid account
func checkCityIsExistByUsingGoogleService(cityName string) bool{
	client, error := maps.NewClient(maps.WithAPIKey(config.GoogleKey))
	if error != nil {
		log.Fatalf("fatal error when create google client: %s", error)
	}
	request := &maps.FindPlaceFromTextRequest{
		Input:     cityName,
		InputType: maps.FindPlaceFromTextInputTypeTextQuery,
	}
	response, err := client.FindPlaceFromText(context.Background(), request)
	if err != nil {
		log.Fatalf("fatal error when make findPlacceFromText request: %s", error)
	}
	return len(response.Candidates) > 0
}
