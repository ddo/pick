package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ddo/pick"
)

// if you pick twice by 1 reader we need to clone that reader.
// since we cant rewind the reader
func main() {
	res, _ := http.Get("http://ddo.me/")
	defer res.Body.Close()

	// create done channel to sync
	done := make(chan bool)

	// clone reader
	attrR, attrW := io.Pipe()
	textR, textW := io.Pipe()

	go func() {
		defer attrW.Close()
		defer textW.Close()

		multiW := io.MultiWriter(attrW, textW)

		// copy the data into the multiwriter
		_, err := io.Copy(multiW, res.Body)

		if err != nil {
			panic(err)
		}
	}()

	go func() {
		href := pick.PickAttr(&pick.Option{
			PageSource: attrR,
			TagName:    "a",
			Attr:       nil,
		}, "href", 0) // limit < 1 = no limit

		fmt.Println("href", href)

		done <- true
	}()

	go func() {
		jumbotron := pick.PickText(&pick.Option{
			PageSource: textR,
			TagName:    "div",
			Attr: &pick.Attr{
				"class",
				"jumbotron",
			},
		})

		fmt.Println("jumbotron", jumbotron)

		done <- true
	}()

	<-done
	<-done
}
