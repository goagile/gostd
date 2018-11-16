package interpret

import (
	"testing"
)

func Test_Each_One(t *testing.T) {
	want := "A"

	data := map[string]interface{}{
		"persons": []interface{}{
			map[string]interface{}{
				"first_name": "A",
			},
		},
	}

	each := GetEach("persons",
		Get("first_name"))

	result := each.From(data).([]interface{})

	if len(result) != 1 {
		t.Error("want len 1")
	} else {
		item := result[0].(map[string]interface{})
		got := item["first_name"]
		if want != got {
			t.Errorf("\n want: %v\n got : %v\n", want, got)
		}
	}
}

func Test_Each_Two(t *testing.T) {
	data := map[string]interface{}{
		"persons": []interface{}{
			map[string]interface{}{
				"first_name": "A",
				"last_name":  "B",
			},
		},
	}

	each := GetEach("persons",
		Get("first_name"),
		Get("last_name"))

	result := each.From(data).([]interface{})

	if len(result) != 1 {
		t.Error("want len 1")
	} else {
		item := result[0].(map[string]interface{})

		want := "A"
		got := item["first_name"]
		if want != got {
			t.Errorf("\n want: %v\n got : %v\n", want, got)
		}

		want = "B"
		got = item["last_name"]
		if want != got {
			t.Errorf("\n want: %v\n got : %v\n", want, got)
		}
	}
}
