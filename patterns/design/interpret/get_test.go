package interpret

import (
	"testing"
)

func Test_Get_One_Field(t *testing.T) {
	want := "A"

	data := map[string]interface{}{
		"fname": "A",
	}

	get := Get("fname")

	got := get.From(data)

	if want != got {
		t.Errorf("\n want: %v\n got : %v\n", want, got)
	}
}

func Test_Get_Two_Fields(t *testing.T) {
	want := "A"

	data := map[string]interface{}{
		"person": map[string]interface{}{
			"fname": "A",
		},
	}

	get := Get("person",
		Get("fname"))

	got := get.From(data)

	if want != got {
		t.Errorf("\n want: %v\n got : %v\n", want, got)
	}
}

func Test_Get_Two_Siblings(t *testing.T) {
	data := map[string]interface{}{
		"person": map[string]interface{}{
			"fname": "A",
			"lname": "B",
		},
	}

	get := Get("person",
		Get("fname"),
		Get("lname"))

	rawresult := get.From(data)

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
