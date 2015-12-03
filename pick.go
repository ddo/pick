package pick

import (
	"io"

	"golang.org/x/net/html"
)

type Option struct {
	PageSource io.Reader
	TagName    string
	Attr       *Attr // optional
}

type Attr struct {
	Label string
	Value string
}

func PickAttr(option *Option, AttrLabel string, limit int) (res []string) {
	if option == nil || option.PageSource == nil {
		return
	}

	z := html.NewTokenizer(option.PageSource)

	for {
		tokenType := z.Next()

		switch tokenType {

		// ignore the error token
		// quit on eof
		case html.ErrorToken:
			if z.Err() == io.EOF {
				return
			}

		case html.SelfClosingTagToken:
			fallthrough
		case html.StartTagToken:
			tagName, attr := z.TagName()

			if string(tagName) != option.TagName {
				continue
			}

			var label, value []byte

			matched := false
			tmpRes := []string{}

			// get attr
			for attr {
				label, value, attr = z.TagAttr()

				labelStr := string(label)
				valueStr := string(value)

				// check the attr
				if option.Attr == nil || (option.Attr.Label == labelStr && option.Attr.Value == valueStr) {
					matched = true
				}

				// get the result - even the matched false or true
				if labelStr == AttrLabel {
					tmpRes = append(tmpRes, valueStr)
				}
			}

			// skip the non matched one
			if !matched {
				continue
			}

			// send the result for matched only
			res = append(res, tmpRes...)

			// return when limit
			if limit > 0 && len(res) >= limit {
				return
			}
		}
	}

	return
}

func PickText(option *Option) (res []string) {
	if option == nil || option.PageSource == nil {
		return
	}

	z := html.NewTokenizer(option.PageSource)

	depth := 0

	for {
		tokenType := z.Next()

		switch tokenType {

		// ignore the error token
		// quit on eof
		case html.ErrorToken:
			if z.Err() == io.EOF {
				return
			}

		case html.TextToken:
			if depth > 0 {
				res = append(res, string(z.Text()))
			}

		case html.EndTagToken:
			if depth > 0 {
				depth--
			}

		case html.StartTagToken:
			if depth > 0 {
				depth++
				continue
			}

			tagName, attr := z.TagName()

			if string(tagName) != option.TagName {
				continue
			}

			var label, value []byte

			matched := false

			// get attr
			for attr {
				label, value, attr = z.TagAttr()

				// TODO: break when found
				if option.Attr == nil || (option.Attr.Label == string(label) && option.Attr.Value == string(value)) {
					matched = true
				}
			}

			if !matched {
				continue
			}

			depth++
		}
	}

	return
}
