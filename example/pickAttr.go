package main

import (
	"fmt"
	"net/http"

	"github.com/ddo/pick"
)

func main() {
	res, _ := http.Get("http://ddo.me/")
	defer res.Body.Close()

	href := pick.PickAttr(&pick.Option{
		PageSource: res.Body,
		TagName:    "a",
		Attr:       nil,
	}, "href", 0) // limit < 1 = no limit

	fmt.Println("href", href)
}
