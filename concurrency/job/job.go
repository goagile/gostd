package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID   int
	Data int
}

type Result struct {
	JobID int
	Data  int
	Time  time.Time
}

var (
	jobs = make(chan Job)
	results = make(chan Result)
)

func main() {
	start := time.Now()



	data := []int{2, 4, 3, 1, 5}

	go allocatejobs(data)

	go makeworks(5)

	for r := range results {
		fmt.Printf("Result{JobID:%2v, Data:%2v, Time:%10v}\n", 
			r.JobID, 
			r.Data, 
			time.Now().Sub(r.Time),
		)
	}



	fmt.Println("Total time:", time.Now().Sub(start))
}

func allocatejobs(data []int) {
	for id, d := range data {
		jobs <- Job{id+1, d}
	}
	close(jobs)
}

func makeworks(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work()
		}()
	}
	wg.Wait()
	close(results)
}

func work() {
	for job := range jobs {
		results <- Result{job.ID, sqrt(job.Data), time.Now()}
	}
}

func sqrt(x int) int {
	time.Sleep(time.Duration(x) * time.Second)
	return x * x
}
