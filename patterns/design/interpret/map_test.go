package main

import "testing"

func Test_X(t *testing.T) {

	data := map[string]interface{}{}

	got := m.Map(data)

	t.Error(got)
}

//
// Mapper
//
// func Mapper(rule Rule) *mapper {
// 	return &mapper{rule}
// }

// type mapper struct {
// 	rule Rule
// }

// func (m *mapper) Map(data interface{}) map[string]interface{} {
// 	result := map[string]interface{}{}
// 	m.rule.Apply(result, data)
// 	return result
// }

//
// Setter
//
type Setter interface {
	Set(key string, value interface{})
}

//
// Getter
//
type Getter interface {
	Get(key string)
}

//
// Set
//
func Set(key string) *set {
	return &set{key}
}

type set struct {
	key string
}

func (s *set) Set(key string, data interface{}) {
}
