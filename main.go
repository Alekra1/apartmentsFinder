package main

import(
	"fmt"
	"github.com/gocolly/colly"
)

type Dictionary map[string]string

type Apartament struct {
	url, name string
	price int

}

func main() {
	//var apartaments []Apartament

	//domain := "https://www.alo.bg"

	//var test string

	url := "https://www.alo.bg/obiavi/imoti-naemi/apartamenti/?region_id=1&order_by=price-asc&io=1"
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
	//collector.OnHTML(".listvip-item-title", func(e *colly.HTMLElement) {
	//	//fmt.Println("found", "content_container_list")
	//	//apartament := Apartament{}
	//	//apartaments_dictionary := Dictionary{}
	//	fmt.Println(e.Text)
	//
	//})
	startingPoint := "javascript:order_by('14')"
	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//fmt.Println("found", "content_container_list")
		//apartament := Apartament{}
		//apartaments_dictionary := Dictionary{}
		if e.Attr("href") == startingPoint{

		}
		//fmt.Println(e.Attr("title"))
		fmt.Println(e.Attr("href"))
		//e.ForEach("a[href]", func(i int, element *colly.HTMLElement) {
		//	link := e.Attr("href")
		//	e.Request.Visit(link)
		//})

	})

	collector.Visit(url)
}
