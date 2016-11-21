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
	}, 0)

	if !reflect.DeepEqual(res, []string{"some text here", "some text here also"}) {
		t.Fail()
	}
}

func TestPickTextLimit(t *testing.T) {
	pageSource := "<div>notme<p>should not include me</p>notme<p class='target'>some text here</p><p class='target'>some text here also</p>notme</div>"

	res := PickText(&Option{
		PageSource: strings.NewReader(pageSource),
		TagName:    "p",
		Attr: &Attr{
			"class",
			"target",
		},
	}, 1)

	if !reflect.DeepEqual(res, []string{"some text here"}) {
		t.Fail()
	}
}

func TestPickTextTree(t *testing.T) {
	pageSource := "<div class='target'><script>console.log('haha')</script><div><p>text1</p>text2<ul><li>text3</li><li>text4</li></ul></div></div>"

	res := PickText(&Option{
		PageSource: strings.NewReader(pageSource),
		TagName:    "div",
		Attr: &Attr{
			"class",
			"target",
		},
	}, 0)

	if !reflect.DeepEqual(res, []string{"console.log('haha')text1text2text3text4"}) {
		t.Fail()
	}
}

func TestPickHtml(t *testing.T) {
	pageSource := "<div class='target'><div><p>text1</p>text2<ul><li>text3</li><li>text4</li></ul></div></div><div class='target'><div><p>text1</p>text2<ul><li>text3</li><li>text4</li></ul></div></div>"

	res := PickHtml(&Option{
		PageSource: strings.NewReader(pageSource),
		TagName:    "div",
		Attr: &Attr{
			"class",
			"target",
		},
	}, 0)

	if !reflect.DeepEqual(res, []string{"<div><p>text1</p>text2<ul><li>text3</li><li>text4</li></ul></div>", "<div><p>text1</p>text2<ul><li>text3</li><li>text4</li></ul></div>"}) {
		t.Fail()
	}
}

func TestPickHtmlLimit(t *testing.T) {
	pageSource := "<div class='target'><div><p>text1</p>text2<ul><li>text3</li><li>text4</li></ul></div></div><div class='target'><div><p>text1</p>text2<ul><li>text3</li><li>text4</li></ul></div></div>"

	res := PickHtml(&Option{
		PageSource: strings.NewReader(pageSource),
		TagName:    "div",
		Attr: &Attr{
			"class",
			"target",
		},
	}, 1)

	if !reflect.DeepEqual(res, []string{"<div><p>text1</p>text2<ul><li>text3</li><li>text4</li></ul></div>"}) {
		t.Fail()
	}
}

func TestPickHtmlNilAttr(t *testing.T) {
	pageSource := "<div class='target'><div><input/><p>text1</p>text2<ul><li>text3</li><li>text4</li></ul></div></div><div class='target'><div><p>text1</p>text2<ul><li>text3</li><li>text4</li></ul></div></div>"

	res := PickHtml(&Option{
		PageSource: strings.NewReader(pageSource),
		TagName:    "ul",
	}, 0)

	if !reflect.DeepEqual(res, []string{"<li>text3</li><li>text4</li>", "<li>text3</li><li>text4</li>"}) {
		t.Fail()
	}
}

var selfClosingTag = [][]byte{
	[]byte("input"),
	[]byte("br"),
	[]byte("hr"),
}

func TestIsSelfClosingTag(t *testing.T) {
	for i := 0; i < len(selfClosingTag); i++ {
		if !isSelfClosingTag(selfClosingTag[i]) {
			t.Fail()
		}
	}
}
