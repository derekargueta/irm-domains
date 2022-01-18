package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	file, err := os.OpenFile("scannedtext.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) //       path 3
	if err != nil {
		log.Println(err.Error() + "cant open results in tempirmdata")
	}

	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		file.WriteString(r.Ctx.Get("url") + "\n")
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("website", err.Error(), "not found on", r.Request.URL.String())
	})

	c.Visit("https://wikipedia.org")
}
