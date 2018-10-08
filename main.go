package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewsPost — a model representing post
type NewsPost struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// FetchParseHabr is a method that retrieves and parses posts from habr.com
func FetchParseHabr() []NewsPost {
	res, err := http.Get("https://habr.com")
	if err != nil {
		panic("Couldn't reach habr.com.")
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Couldn't create document from response.")
	}

	var posts []NewsPost

	doc.Find(".post_preview").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".post__title_link").Text()
		text := s.Find(".post__text.post__text-html").Text()
		posts = append(posts, NewsPost{Title: title, Text: text})
	})

	return posts
}

func main() {
	//FetchParseHabr()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/posts", func(c echo.Context) error {
		return c.JSON(http.StatusOK, FetchParseHabr())
	})

	e.File("/", "dist/index.html")
	e.Static("/", "dist")
	e.Start(":1323")
}
