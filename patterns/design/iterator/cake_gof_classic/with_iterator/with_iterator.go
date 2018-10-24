package main

import (
	"fmt"
	"github.com/khardi/trygo/dp/iterator/cake_gof_classic/cake"
)

func main() {
  print("Pancake menu", cake.NewPancakeHouseMenu())
	print("Diner menu", cake.NewDinerMenu())
  print("Ice Cream menu", cake.NewIceCreamMenu())
}

func print(menuName string, menu cake.Iterabler) {
	fmt.Printf("\n%s:\n", menuName)
	fmt.Println("-------------")

	// В этой реализации итератор является локальной переменной цикла
	for it := menu.Iterator(); it.HasNext(); {
    fmt.Println(it.Next())
  }
}
