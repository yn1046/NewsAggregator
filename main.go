package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://habr.com")
	if err != nil {
		panic("sasee")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	println(string(body))
}
