package main

import "fmt"
import "sort"


func main() {

	// Хеш-таблицу можно создать с помощью функции make
	// Таблица создается пустой
	m := make(map[string]int)
	fmt.Println("m:", m)
	m["X"] = 1
	m["Y"] = 2
	fmt.Println("m:", m)

	// Создание Хэш-таблицы со значениями по умолчанию
	z := map[string]int{
		"X": 1,
		"Y": 2,
	}
	fmt.Println("z:", z)

	// Удаление ключа из Хэш-таблицы
	delete(z, "X")
	fmt.Println("z:", z)

	fmt.Println("Итерирование по ключам Хэш-таблицы")
	for k := range m {
		fmt.Printf("%v\n", k)
	}	

	fmt.Println("Итерирование по ключам и значениям Хэш-таблицы")
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}

	fmt.Println("Сортированные ключи и значения:")
	ages := map[string]int{
		"Bob": 23,
		"Alice": 12,
	}
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s: %v\n", name, ages[name])
	}
	
	fmt.Println("Проверка существования ключа:")
	if _, ok := ages["Anit"]; !ok {
		fmt.Println("Не найден ключ Anit")
	}

}
