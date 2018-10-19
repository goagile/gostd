package main

import (
	"fmt"
	"time"
)

func main() {
	// это входные данные
	numbers := []int{1, 2, 3}

	// это очередь входных данных
	data := make(chan int)

	// это очередь выходных данных
	result := make(chan int)

	// создаем горутину для каждого элемента входных данных
	for i := 0; i < len(numbers); i++ {
		go calc(data, result)
	}

	// заполняем очередь входными данными
	for _, x := range numbers {
		data <- x
	}

	// запускаем горутины для чтения
	for i := 0; i < len(numbers); i++ {
		go read(result)
	}

	time.Sleep(2 * time.Second)
}

func calc(data chan int, result chan int) {
	// получаем данные из входной очереди
	x := <-data

	time.Sleep(time.Second * time.Duration(x))

	// вычисляем результат
	square := x * x

	// отправляем данные в выходную очередь
	result <- square
}

func read(result chan int) {
	fmt.Println(<-result)
}
