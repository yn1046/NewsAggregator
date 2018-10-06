package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// FetchParseHabr is a method that retrieves and parses posts from habr.com
func FetchParseHabr() {
	res, err := http.Get("https://habr.com")
	if err != nil {
		panic("Couldn't reach habr.com.")
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Couldn't create document from response.")
	}

	doc.Find(".post_preview").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".post__title_link").Text()
		text := s.Find(".post__text.post__text-html").Text()
		println(title)
		println(text + "\n\n\n")

	})
}

func main() {
	//FetchParseHabr()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "client/index.html")

}
