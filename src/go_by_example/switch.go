package main

import (
	"fmt"
	// "time"
)

func main() {

	// fmt.Println("---- Switch ----")
	// i := 2
	// fmt.Print("Write ", i, " as ")
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// }

	// fmt.Println("---- Weekday ----")
	// w := time.Now().Weekday()
	// switch w {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It is the weekend")
	// default:
	// 	fmt.Println("It is a weekday")
	// }

	// fmt.Println("---- Time ----")
	// t := time.Now()
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("It is before noon")
	// default: 
	// 	fmt.Println("It is after noon")
	// }

	fmt.Println("---- What Am I ----")
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case float64:
			fmt.Println("float64")
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Println(t)
		}
	}

	whatAmI(true)
	whatAmI(1.0)
	whatAmI(1)
	whatAmI("a")

}
