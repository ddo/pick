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
	pageSource := "<a href='http://ddo.me'>test</a><a href='http://ddict.me'>test</a>"

	a, err := PickAttr(&PickOption{
		&pageSource,
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
	pageSource := "<a href='http://ddo.me'>test</a><a id='target' href='http://ddict.me'>test</a>"

	a, err := PickAttr(&PickOption{
		&pageSource,
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
	pageSource := "<a href='http://ddo.me'>test</a><a id='targett' href='http://ddict.me'>test</a>"

	a, err := PickAttr(&PickOption{
		&pageSource,
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
	pageSource := "<input type='text' id='target' value='haha' />"

	input, err := PickAttr(&PickOption{
		&pageSource,
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

func TestPickText(t *testing.T) {
	pageSource := "<div>notme<p>should not include me</p>notme<p class='target'>some text here</p><p class='target'>some text here also</p>notme</div>"

	data, err := PickText(&PickOption{
		&pageSource,
		"p",
		&Attr{
			"class",
			"target",
		},
	})

	if err != nil {
		t.Fail()
		return
	}

	if !reflect.DeepEqual(data, []string{"some text here", "some text here also"}) {
		t.Fail()
	}
}

func TestPickTextTree(t *testing.T) {
	pageSource := "<div class='target'><div><p>text1</p>text2<ul><li>text3</li><li>text4</li></ul></div></div>"

	data, err := PickText(&PickOption{
		&pageSource,
		"div",
		&Attr{
			"class",
			"target",
		},
	})

	if err != nil {
		t.Fail()
		return
	}

	if !reflect.DeepEqual(data, []string{"text1", "text2", "text3", "text4"}) {
		t.Fail()
	}
}
