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
		tokenType := z.Next()

		switch tokenType {

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

			data_tmp := []string{}

			matched := false

			//get attr
			for attr {
				label, value, attr = z.TagAttr()

				label_str := string(label)
				value_str := string(value)

				if option.Attr == nil || (option.Attr.Label == label_str && option.Attr.Value == value_str) {
					matched = true
				}

				if label_str == AttrLabel {
					data_tmp = append(data_tmp, value_str)
				}
			}

			if !matched {
				continue
			}

			//loop attr
			data = append(data, data_tmp...)
		}
	}

	return data, z.Err()
}
