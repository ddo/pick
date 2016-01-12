package pick

import (
	"reflect"
	"strings"
	"testing"
)

func TestPickEmptyOption(t *testing.T) {
	res := PickAttr(nil, "href", 0)

	if len(res) > 0 {
		t.Fail()
	}
}

func TestPickAttrEmptyAttrOption(t *testing.T) {
	pageSource := "<a href='http://ddo.me'>test</a><a href='http://ddict.me'>test</a>"

	res := PickAttr(&Option{
		PageSource: strings.NewReader(pageSource),
		TagName:    "a",
		Attr:       nil,
	}, "href", 0)

	if !reflect.DeepEqual(res, []string{"http://ddo.me", "http://ddict.me"}) {
		t.Fail()
	}
}

func TestPickAttr(t *testing.T) {
	pageSource := "<a href='http://ddo.me'>test</a><a id='target' href='http://ddict.me'>test</a>"

	res := PickAttr(&Option{
		PageSource: strings.NewReader(pageSource),
		TagName:    "a",
		Attr: &Attr{
			"id",
			"target",
		},
	}, "href", 0)

	if !reflect.DeepEqual(res, []string{"http://ddict.me"}) {
		t.Fail()
	}
}

func TestPickAttrLimit(t *testing.T) {
	pageSource := "<a href='http://ddo.me'>test</a><a href='http://ddict.me'>test</a>"

	res := PickAttr(&Option{
		PageSource: strings.NewReader(pageSource),
		TagName:    "a",
		Attr:       nil,
	}, "href", 1)

	if !reflect.DeepEqual(res, []string{"http://ddo.me"}) {
		t.Fail()
	}
}

func TestPickAttrFail(t *testing.T) {
	pageSource := "<a href='http://ddo.me'>test</a><a id='targett' href='http://ddict.me'>test</a>"

	res := PickAttr(&Option{
		PageSource: strings.NewReader(pageSource),
		TagName:    "a",
		Attr: &Attr{
			"id",
			"target",
		},
	}, "href", 0)

	if len(res) > 0 {
		t.Fail()
	}
}

func TestPickAttrSelfClosingTagToken(t *testing.T) {
	pageSource := "<input type='text' id='target' value='haha' />"

	res := PickAttr(&Option{
		PageSource: strings.NewReader(pageSource),
		TagName:    "input",
		Attr: &Attr{
			"id",
			"target",
		},
	}, "value", 0)

	if !reflect.DeepEqual(res, []string{"haha"}) {
		t.Fail()
	}
}

func TestPickText(t *testing.T) {
	pageSource := "<div>notme<p>should not include me</p>notme<p class='target'>some text here</p><p class='target'>some text here also</p>notme</div>"

	res := PickText(&Option{
		PageSource: strings.NewReader(pageSource),
		TagName:    "p",
		Attr: &Attr{
			"class",
			"target",
		},
	})

	if !reflect.DeepEqual(res, []string{"some text here", "some text here also"}) {
		t.Fail()
	}
}

func TestPickTextTree(t *testing.T) {
	pageSource := "<div class='target'><div><p>text1</p>text2<ul><li>text3</li><li>text4</li></ul></div></div>"

	res := PickText(&Option{
		PageSource: strings.NewReader(pageSource),
		TagName:    "div",
		Attr: &Attr{
			"class",
			"target",
		},
	})

	if !reflect.DeepEqual(res, []string{"text1", "text2", "text3", "text4"}) {
		t.Fail()
	}
}
