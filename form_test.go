package pick

import (
	"reflect"
	"strings"
	"testing"
)

const (
	form_testcase = `
<!DOCTYPE html>
<html>
<body>

<form id="login_form" action="action_page.php">
    First name:
    <input type="text" name="email" value="email@ddo.me">

    Last name:
    <input type="text" name="password" value="password">

    Country:
    <select name="country">
        <option value="usa">usa</option>
        <option value="canada" selected>canada</option>
    </select>

    Age:
    <select name="age">
        <option value=1>1</option>
        <option value=2>2</option>
        <option value=3>3</option>
        <option value=4>4</option>
    </select>

    Description:
    <textarea name="description">description details</textarea>

    <input name="submit" type="submit" value="login">
</form> 

<p>Forget password? us too lol</p>

</body>
</html>
    `
)

func TestPickForm(t *testing.T) {
	input := PickForm(strings.NewReader(form_testcase), &Attr{"id", "login_form"})
	if input == nil {
		t.Error()
		return
	}

	if !reflect.DeepEqual(input, map[string][]string{
		"email":       []string{"email@ddo.me"},
		"password":    []string{"password"},
		"description": []string{"description details"},
		"country":     []string{"canada"},
		"age":         []string{"1"},
		"submit":      []string{"login"},
	}) {
		t.Error()
	}
}

func TestPickFormEmpty(t *testing.T) {
	input := PickForm(strings.NewReader(form_testcase), &Attr{"id", "form"})
	if input != nil {
		t.Error()
		return
	}
}
