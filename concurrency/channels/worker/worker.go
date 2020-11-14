package main

import "fmt"
import "time"

func main() {

	jobs := make(chan string)
	results := make(chan int)

	go worker(1, jobs, results)

	jobs <- "\tA"
	jobs <- "\t\tB"
	jobs <- "\t\t\tC"
	jobs <- "\t\t\t\tD"
	jobs <- "\t\t\t\t\tE"
	jobs <- "\t\t\t\t\t\tF"
	close(jobs)

	for r := range results {
		fmt.Println(r)
	}

}

func worker(id int, jobs <-chan string, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "job", j, "start ... ")
		time.Sleep(time.Second * time.Duration(id))
		fmt.Println("worker", id, "job", j, "finish")
		results <- id
	}
}
