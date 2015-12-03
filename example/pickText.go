package main

import (
	"fmt"
	"net/http"

	"github.com/ddo/pick"
)

func main() {
	res, _ := http.Get("http://ddo.me/")
	defer res.Body.Close()

	jumbotron := pick.PickText(&pick.Option{
		PageSource: res.Body,
		TagName:    "div",
		Attr: &pick.Attr{
			"class",
			"jumbotron",
		},
	})

	fmt.Println("jumbotron", jumbotron)
}
