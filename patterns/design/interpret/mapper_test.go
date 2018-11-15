package interpret

import (
	"testing"
)

//
// Mapper
//
func Test_Mapper_One(t *testing.T) {
	want := "A"

	mapper := Mapper(
		Set("fname",
			Get("person",
				Get("first_name"))))

	result := mapper.Map(map[string]interface{}{
		"person": map[string]interface{}{
			"first_name": "A",
		},
	})

	got := result["fname"]

	if want != got {
		t.Errorf("\n want: %v\n got : %v\n", want, got)
	}
}

func Test_Mapper_Two(t *testing.T) {
	mapper := Mapper(
		Set("fname",
			Get("person",
				Get("first_name"))),

		Set("lname",
			Get("person",
				Get("last_name"))))

	result := mapper.Map(map[string]interface{}{
		"person": map[string]interface{}{
			"first_name": "A",
			"last_name":  "B",
		},
	})

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

func Test_Mapper_Two_Levels(t *testing.T) {

	mapper := Mapper(
		Set("user",
			Set("fname",
				Get("person",
					Get("first_name"))),
			Set("lname",
				Get("person",
					Get("last_name")))))

	result := mapper.Map(map[string]interface{}{
		"person": map[string]interface{}{
			"first_name": "A",
			"last_name":  "B",
		},
	})

	user := result["user"].(map[string]interface{})

	want := "A"
	got := user["fname"]
	if want != got {
		t.Errorf("\n want: %v\n got : %v\n", want, got)
	}

	want = "B"
	got = user["lname"]
	if want != got {
		t.Errorf("\n want: %v\n got : %v\n", want, got)
	}
}

func Test_Mapper_Decompose(t *testing.T) {

	firstName :=
		Get("person",
			Get("first_name"))

	lastName :=
		Get("person",
			Get("last_name"))

	defaultMidName :=
		Set("mname", "C")

	mapper := Mapper(
		Set("user",
			Set("fname", firstName),
			Set("lname", lastName),
			defaultMidName))

	result := mapper.Map(map[string]interface{}{
		"person": map[string]interface{}{
			"first_name": "A",
			"last_name":  "B",
		},
	})

	user := result["user"].(map[string]interface{})

	want := "A"
	got := user["fname"]
	if want != got {
		t.Errorf("\n want: %v\n got : %v\n", want, got)
	}

	want = "B"
	got = user["lname"]
	if want != got {
		t.Errorf("\n want: %v\n got : %v\n", want, got)
	}

	want = "C"
	got = user["mname"]
	if want != got {
		t.Errorf("\n want: %v\n got : %v\n", want, got)
	}
}
