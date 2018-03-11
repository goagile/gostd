package main

import (
	"fmt"
	"github.com/khardi/trygo/dp/iterator/cake_gof_classic/cake"
)

func main() {
  print("Pancake menu", cake.NewPancakeHouseMenu().Iterator())
	print("Diner menu", cake.NewDinerMenu().Iterator())
  print("Ice Cream menu", cake.NewIceCreamMenu().Iterator())
}

func print(menuName string, it cake.Iterator) {
	fmt.Printf("\n%s:\n", menuName)
	fmt.Println("-------------")
	for it.HasNext() {
    fmt.Println(it.Next())
  }
}
