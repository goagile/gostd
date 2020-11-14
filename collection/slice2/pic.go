package main

import (
	"golang.org/x/tour/pic"
)

func main() {
	pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
	rows := make([][]uint8, dy)
	for y := range rows {
		row := make([]uint8, dx)
		for x := range row {
			row[x] = uint8(x * y)
		} 
		rows[y] = row
	}
	return rows	
}
