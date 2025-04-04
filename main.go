package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Apartaments struct {
	url, name []string
}

func main() {
	//var apartaments []Apartament

	//domain := "https://www.alo.bg"

	//var test string

	url := "https://www.alo.bg/obiavi/imoti-naemi/apartamenti/?region_id=1&order_by=time-desc&io=1"
	collector := colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
		// print the url of that request
		fmt.Println("Visiting", r.URL)
	})
	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})
	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("Blimey, an error occurred!:", e)
	})

	apartaments := Apartaments{}

	collector.OnHTML(".listvip-item-header", func(e *colly.HTMLElement) {
		apartaments.name = e.ChildAttrs("a", "title")
		apartaments.url = e.ChildAttrs("a", "href")

		fmt.Println(apartaments.name)
		fmt.Println(apartaments.url)

	})

	collector.Visit(url)
}
