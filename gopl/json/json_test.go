package main

import (
	// "fmt"
	"encoding/json"
	"testing"
)

func testJSON(t *testing.T) {
	want := map[string]interface{}{
		"x": 3,
	}

	var got map[string]interface{}
	bytes := []byte(`{"x": 2}`)
	err := json.Unmarshal(bytes, &got)
	if err != nil {
		t.Error("Some error: ", err)
	}

	if type want != type got {
		t.Errorf("want != got: %v != %v\n", want, got)
	}

	// if want["x"] != got["x"] {
	// 	t.Errorf("want != got: %v != %v\n", want, got)
	// }

	// if got["x"] != want["x"] {
	// 	t.Errorf("want != got: %v != %v\n", want, got)
	// }

}
