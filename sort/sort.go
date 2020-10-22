package main

import (
	"fmt"
	"sort"
)

func main() {

	i := []int{4, 3, 2, 1, 0}
	sort.Ints(i)
	fmt.Println(i)

	s := []string{"z", "c", "a", "x", "b"}
	sort.Strings(s)
	fmt.Println(s)

	f := []float64{0.56, 0.45, 0.33, 0.1}
	sort.Float64s(f)
	fmt.Println(f)

	r := []int{2, 3, 4, 5}
	sort.Sort(sort.Reverse(sort.IntSlice(r)))
	fmt.Println(r)

	runes := []rune{'g', 'o', 'o', 'g', 'l', 'e'}
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	for _, r := range runes {
		fmt.Printf("%v ", string(r))
	}
	fmt.Println()

	students := []*Student{&Student{21}, &Student{19}, &Student{20}}
	sort.Sort(ByAge(students))
	fmt.Println(students)
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student(%v)", s.Age)
}

type ByAge []*Student

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Len() int {
	return len(a)
}
