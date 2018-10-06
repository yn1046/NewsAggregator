package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://habr.com")
	if err != nil {
		panic("sasee")
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("SHYTE SUXX")
	}

	doc.Find(".post_preview").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".post__title_link").Text()
		text := s.Find(".post__text.post__text-html").Text()
		println(title)
		println(text)
		println("post found!")

	})
}
