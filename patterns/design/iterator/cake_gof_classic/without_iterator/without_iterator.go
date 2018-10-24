package main

import (
	"fmt"
	"github.com/khardi/trygo/dp/iterator/cake_gof_classic/cake"
)

func main() {
	printPancake()
	printDiner()
	printIcecream()
}

func printPancake() {
	fmt.Println("\nPancake menu:")
	fmt.Println("-------------")
	menu := cake.NewPancakeHouseMenu()
	items := menu.MenuItems
	for _, item := range items {
		fmt.Println(item)
	}
}

func printDiner() {
	fmt.Println("\nDiner menu:")
	fmt.Println("-------------")
	menu := cake.NewDinerMenu()
	items := menu.MenuItems
	for i := 0; i < len(items); i++ {
		fmt.Println(items[i])
	}
}

func printIcecream() {
	fmt.Println("\nIce Cream menu:")
	fmt.Println("-------------")
	menu := cake.NewIceCreamMenu()
	items := menu.MenuItems
	for _, item := range items {
		fmt.Println(item)
	}
}
