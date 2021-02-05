package cities_api

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"strings"
	"sync"
)

var (
	suffixes = []string{"cityA", "cityB", "cityV", "cityG", "cityD", "cityE", "cityZh", "cityZ", "cityI", "cityIi",
		"cityK", "cityL", "cityM", "cityN", "cityO", "cityP", "cityR", "cityS", "cityT", "cityU", "cityF", "cityH",
		"cityTs", "cityCh", "citySh", "citySch", "cityYe", "cityYu", "cityYa"}
)

type ResultCities struct {
	sync.Mutex
	data map[string]string
}

var resultCities = ResultCities{
	Mutex: sync.Mutex{},
	data:  make(map[string]string),
}

func ParseCities() map[string]string {
	fmt.Printf("\nParce cities start\n")
	newGeziyor := geziyor.NewGeziyor(&geziyor.Options{
		StartURLs:   []string{"http://www.1000mest.ru/"},
		ParseFunc:   parse,
		LogDisabled: true,
	})
	newGeziyor.Start()
	fmt.Printf("\nParce cities end\n")
	return resultCities.data
}

func parse(g *geziyor.Geziyor, r *client.Response) {
	for _, suffix := range suffixes {
		g.Get(r.JoinURL(suffix), func(_g *geziyor.Geziyor, _r *client.Response) {
			_r.HTMLDoc.Find("table").Find("tr").Each(func(i int, s *goquery.Selection) {
				find := s.Find("td").First()
				lower := strings.ToLower(find.Text())
				if len(lower) > 0 {
					resultCities.Lock()
					resultCities.data[lower] = lower
					resultCities.Unlock()
				}
			})
		})
	}
}
