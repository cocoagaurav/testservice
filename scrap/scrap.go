package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"time"
)

func Scrap(date string) string {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	//	for _=range time.Tick(5*time.Second) {
	request, err := http.NewRequest("GET", "https://www.eduro.com/", nil)
	if err != nil {
		log.Printf("unable to create request due to error:=[%v]", err)
	}
	request.Header.Set("User-Agent", "Not Firefox")
	response, err := client.Do(request)
	if err != nil {
		log.Printf("unable to get response due to error:=[%v]", err)

	}
	return ScrapQuote(response, date)
	//	}

}

func ScrapQuote(response *http.Response, querydate string) string {

	document, err := goquery.NewDocumentFromReader(response.Body)
	var resp string
	if err != nil {
		log.Printf("unable to create reader due to err:[%v]", err)
	}
	document.Find("div").Each(func(index int, element *goquery.Selection) {
		class, exist := element.Attr("class")
		if exist {
			if class == "datebox" {
				date := element.Children().Contents().Text()
				if date == querydate {
					element.Parent().Parent().Parent().Find("div").EachWithBreak(func(index int, ele *goquery.Selection) bool {
						class1, exist := ele.Attr("class")
						if exist {
							if class1 == "article" {
								resp = ele.Children().Children().Contents().Text()
								return false
							}
						}
						return true
					})
				}
			}
		}
	})
	return resp
}
