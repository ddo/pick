package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ddo/pick"
)

func main() {
	res, _ := http.Get("http://ddo.me/")
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	html := string(body)

	a, _ := pick.PickAttr(&pick.PickOption{
		html,
		"a",
		nil,
	}, "href")

	fmt.Println(a)

	jumbotron, _ := pick.PickText(&pick.PickOption{
		html,
		"div",
		&pick.Attr{
			"class",
			"jumbotron",
		},
	})

	fmt.Println(jumbotron)
}
