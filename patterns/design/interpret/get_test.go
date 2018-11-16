package interpret

import (
	"testing"
)

func Test_Get_One_Field(t *testing.T) {
	want := "A"

	inputresult := map[string]interface{}{}

	data := map[string]interface{}{
		"fname": "A",
	}

	get := Get("fname")

	got := get.From(inputresult, data)

	if want != got {
		t.Errorf("\n want: %v\n got : %v\n", want, got)
	}
}

func Test_Get_Two_Fields(t *testing.T) {
	want := "A"

	inputresult := map[string]interface{}{}

	data := map[string]interface{}{
		"person": map[string]interface{}{
			"fname": "A",
		},
	}

	get := Get("person",
		Get("fname"))

	got := get.From(inputresult, data)

	if want != got {
		t.Errorf("\n want: %v\n got : %v\n", want, got)
	}
}

func Test_Get_Two_Siblings(t *testing.T) {

	inputresult := map[string]interface{}{}

	data := map[string]interface{}{
		"person": map[string]interface{}{
			"fname": "A",
			"lname": "B",
		},
	}

	get := Get("person",
		Get("fname"),
		Get("lname"))

	rawresult := get.From(inputresult, data)

	result := rawresult.(map[string]interface{})

	want := "A"
	got := result["fname"]
	if want != got {
		t.Errorf("\n want: %v\n got : %v\n", want, got)
	}

	want = "B"
	got = result["lname"]
	if want != got {
		t.Errorf("\n want: %v\n got : %v\n", want, got)
	}
}

func Test_Get_Set_One(t *testing.T) {
	// want := "A"

	inputresult := map[string]interface{}{}

	data := map[string]interface{}{
		"fname": "A",
	}

	get :=
		Get("fname",
			Set("first_name"))

	get.From(inputresult, data)

	t.Error(inputresult)

	// if want != got {
	// 	t.Errorf("\n want: %v\n got : %v\n", want, got)
	// }
}
