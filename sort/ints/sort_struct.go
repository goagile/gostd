package main

import (
	"fmt"
	"sort"
)


type Color struct {
	Area float64
}


type ByArea []Color
func (c ByArea) Len() int {
	return len(c)
}
func (c ByArea) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c ByArea) Less(i, j int) bool {
	return c[i].Area > c[j].Area
}


func main() {

	colors := []Color{
		{5.32}, 
		{1.51},
		{3.31},
	}

	fmt.Println("Not sorted colors")
	fmt.Println(colors)

	sort.Sort(ByArea(colors))
	
	fmt.Println("Sorted colors")
	fmt.Println(colors)	

}
