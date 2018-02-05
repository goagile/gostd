package main

import "fmt"

func main() {

	months := []string{
		1: "Январь",
		2: "Февраль",
		3: "Март",
		4: "Апрель",
		5: "Май",
		6: "Июнь",
		7: "Июль",
		8: "Август",
		9: "Сентябрь",
		10: "Октябрь",
		11: "Ноябрь",
		12: "Декабрь",
	}

	fmt.Println("\nПоследний месяц года")
	fmt.Println(months[len(months)-1])

	fmt.Println("\nПеречислим все месяцы")
	for i, month := range months {
		fmt.Printf("months[%v] = %q\n", i, month)
	}

	fmt.Println("\nМесяцы лета:")
	summer := months[6:9]
	for _, m := range summer {
		fmt.Println(m)	
	}

}
