package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

func scrap(browser string) (ua []string, err error) {
	c := colly.NewCollector(colly.AllowedDomains("useragentstring.com"))

	c.OnHTML("body", func(h *colly.HTMLElement) {
		h.ForEach("ul li", func(i int, h *colly.HTMLElement) {
			ua = append(ua, h.Text)
		})
	})

	err = c.Visit(fmt.Sprintf("https://useragentstring.com/pages/%s/", browser))
	if err != nil {
		return
	}

	return
}

func main() {
	browsers := []string{
		"chrome",
		"edge",
		"firefox",
		"safari",
	}

	result := map[string]interface{}{}

	output := flag.String("output", "", "output file location (json)")
	flag.Parse()

	if *output == "" {
		flag.Usage()
		return
	}

	for _, browser := range browsers {
		ua, err := scrap(browser)
		if err != nil {
			log.Fatal(err)
		}

		result[browser] = ua
	}

	f, err := os.OpenFile(*output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(f).Encode(result)
	if err != nil {
		log.Fatal(err)
	}
}
