package interpret

import (
	"testing"
)

func Test_Set_One_Field(t *testing.T) {
	want := "A"

	result := map[string]interface{}{}

	set := Set("fname", "A")

	set.To(result, map[string]interface{}{})

	got := result["fname"]

	if want != got {
		t.Errorf("\n want: %v\n got : %v\n", want, got)
	}
}

func Test_Set_Two_Fields(t *testing.T) {
	want := "A"

	result := map[string]interface{}{}

	set :=
		Set("person",
			Set("fname", "A"))

	set.To(result, map[string]interface{}{})

	person := result["person"].(map[string]interface{})

	got := person["fname"]

	if want != got {
		t.Errorf("\n want: %v\n got : %v\n", want, got)
	}
}

func Test_Set_Two_Fields_Inherit(t *testing.T) {
	result := map[string]interface{}{}

	set :=
		Set("person",
			Set("fname", "A"),
			Set("lname", "B"))

	set.To(result, map[string]interface{}{})

	person := result["person"].(map[string]interface{})

	want := "A"
	got := person["fname"]
	if want != got {
		t.Errorf("\n want: %v\n got : %v\n", want, got)
	}

	want = "B"
	got = person["lname"]
	if want != got {
		t.Errorf("\n want: %v\n got : %v\n", want, got)
	}
}
