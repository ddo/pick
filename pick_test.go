package pick

import (
	// "fmt"
	"reflect"
	"testing"
)

func TestPickEmptyOption(t *testing.T) {
	a, err := PickAttr(nil, "href")

	if err != nil {
		t.Fail()
		return
	}

	if len(a) > 0 {
		t.Fail()
	}
}

func TestPickAttrEmptyAttrOption(t *testing.T) {
	a, err := PickAttr(&PickOption{
		"<a href='http://ddo.me'>test</a><a href='http://ddict.me'>test</a>",
		"a",
		nil,
	}, "href")

	if err != nil {
		t.Fail()
		return
	}

	if !reflect.DeepEqual(a, []string{"http://ddo.me", "http://ddict.me"}) {
		t.Fail()
	}
}

func TestPickAttr(t *testing.T) {
	a, err := PickAttr(&PickOption{
		"<a href='http://ddo.me'>test</a><a id='target' href='http://ddict.me'>test</a>",
		"a",
		&Attr{
			"id",
			"target",
		},
	}, "href")

	if err != nil {
		t.Fail()
		return
	}

	if !reflect.DeepEqual(a, []string{"http://ddict.me"}) {
		t.Fail()
	}
}

func TestPickAttrFail(t *testing.T) {
	a, err := PickAttr(&PickOption{
		"<a href='http://ddo.me'>test</a><a id='targett' href='http://ddict.me'>test</a>",
		"a",
		&Attr{
			"id",
			"target",
		},
	}, "href")

	if err != nil {
		t.Fail()
		return
	}

	if len(a) > 0 {
		t.Fail()
	}
}

func TestPickAttrSelfClosingTagToken(t *testing.T) {
	input, err := PickAttr(&PickOption{
		"<input type='text' id='target' value='haha' />",
		"input",
		&Attr{
			"id",
			"target",
		},
	}, "value")

	if err != nil {
		t.Fail()
		return
	}

	if !reflect.DeepEqual(input, []string{"haha"}) {
		t.Fail()
	}
}
