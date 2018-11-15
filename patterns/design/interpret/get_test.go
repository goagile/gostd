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
