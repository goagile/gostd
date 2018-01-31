package main

import (
	"fmt"
	"sort"
)

func main() {

	fruits := []string{"абрикосович", "яблок", "бан"}
	sort.Sort(ByLen(fruits))
	fmt.Println(fruits)

}

type ByLen []string
func (s ByLen) Len() int {
	return len(s)
}
func (s ByLen) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLen) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
