package main

import "testing"

func Test_Map(t *testing.T) {

	got := Map(Identity, []interface{}{1, 2, 3})

	for _, g := range got {
		t.Error(g)
	}
}

func Map(f func(interface{}) interface{}, slice []interface{}) []interface{} {
	result := []interface{}{}
	for _, x := range slice {
		result = append(result, f(x))
	}
	return result
}

func Identity(raw interface{}) (result interface{}) {
	z := map[string]interface{}{}
	id := raw.(int)
	z["id"] = id
	result = z
	return
}
