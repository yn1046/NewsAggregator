package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	response, err := http.Get("https://habr.com")
	if err != nil {
		panic("sasee")
	}
	z := html.NewTokenizer(response.Body)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			isAnchor := t.Data == "li"
			if isAnchor {
				fmt.Println("We found a link!")
				fmt.Println("It has attribute" + t.Attr[0].Val)
			}
		}
	}
}
