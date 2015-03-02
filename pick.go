package pick

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type PickOption struct {
	PageSource string
	TagName    string
	Attr       *Attr //optional
}

type Attr struct {
	Label string
	Value string
}

func PickAttr(option *PickOption, AttrLabel string) (data []string, err error) {
	if option == nil || option.PageSource == "" {
		return data, nil
	}

	z := html.NewTokenizer(strings.NewReader(option.PageSource))

	for {
		tt := z.Next()

		switch tt {

		//ignore the error token
		//quit on eof
		case html.ErrorToken:
			if z.Err() == io.EOF {
				return data, nil
			}

		case html.SelfClosingTagToken:
			fallthrough
		case html.StartTagToken:
			tagName, attr := z.TagName()

			if string(tagName) != option.TagName {
				continue
			}

			var label, value []byte

			attr_arr := []*Attr{}

			matched := false

			//get attr
			for attr {
				label, value, attr = z.TagAttr()

				label_str := string(label)
				value_str := string(value)

				if option.Attr == nil || (option.Attr.Label == label_str && option.Attr.Value == value_str) {
					matched = true
				}

				attr_arr = append(attr_arr, &Attr{
					label_str,
					value_str,
				})
			}

			if !matched {
				continue
			}

			//loop attr
			for i := 0; i < len(attr_arr); i++ {
				attr := attr_arr[i]

				if attr.Label == AttrLabel {
					data = append(data, attr.Value)
				}
			}
		}
	}

	return data, z.Err()
}
