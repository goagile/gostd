package main

import "fmt"
import "time"

func main() {

	// создаем каналы работ и результатов
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// запускаем 3 worker'a
	// каждому worker'у даем id, каналы работ и результатов
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// наполняем канал работ
	for j := 1; j <= 5; j++ {
		jobs<- j
	}
	// закрываем наполненный канал работ
	close(jobs)

	// читаем канал результатов
	for a := 1; a <= 5; a++ {
		<-results
	}

}

func worker(id int, jobs <-chan int, results chan<- int) {
	
	// воркер перебирает все работы по очереди
	for j := range jobs {
		// выполняет работу
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second * 1)
		fmt.Println("\tworker", id, "finished job", j)
		// отправляет результат работы в канал результатов
		results <- j * 2
	}

}
