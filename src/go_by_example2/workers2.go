package main

import (
	"fmt"
	"time"
)

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for r := 1; r <= 5; r++ {
		<- results
	}

}


func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "start job", j)
		time.Sleep(time.Second * 2)
		fmt.Println("\tworker", id, "finish job", j)
		result := j * 2
		results <- result
	}
}
