package main

import (
	"fmt"
	"strings"
)

func main() {

	// Строковый массив
	var strs = []string{"apple", "banana", "peach", "plum"}
	fmt.Println("strs:")
	fmt.Println(strs)
	
	// Порядковый номер элемента в массива
	fmt.Println("Index:")
	fmt.Println(Index(strs, "apple"))
	fmt.Println(Index(strs, "banana"))
	fmt.Println(Index(strs, "peach"))
	fmt.Println(Index(strs, "plum"))

	// Включение элемента
	fmt.Println("Include:")
	fmt.Println(Include(strs, "plum"))
	fmt.Println(Include(strs, "arbuz"))

	// Верно если хотя бы один элемент удовлетворяет условию
	fmt.Println("Any HasPrefix p:")
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	// Верно только если все элементы удовлетворяют условию
	fmt.Println("All HasPrefix p:")
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	// Фильтрация значений по условию
	fmt.Println("Filter all Contains e:")
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))


	// Map
	fmt.Println("Map to Upper:")
	fmt.Println(Map(strs, strings.ToUpper))
}

func Index(array []string, v string) int {
	for i, a := range array {
		if a == v {
			return i
		}
	}
	return -1
}

func Include(array []string, v string) bool {
	return Index(array, v) >= 0
}

func Any(array []string, f func(string) bool) bool {
	for _, v := range array {
		if f(v) {
			return true
		}
	}
	return false
}

func All(array []string, f func(string) bool) bool {
	for _, v := range array {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(array []string, f func(string) bool) []string {
	result := make([]string, 0)
	for _, v := range array {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map(array []string, f func(string) string) []string {
	result := make([]string, len(array))
	for i, a := range array {
		result[i] = f(a)
	}
	return result
}
