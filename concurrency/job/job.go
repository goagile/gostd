package main

import (
	"fmt"
	"sync"
	"time"
)

const n = 10

var jobs = make(chan Job, n)

var results = make(chan Result, n)

func main() {
	start := time.Now()

	njobs := 10
	go allocate(njobs)

	done := make(chan bool)
	go result(done)

	nworkers := 10
	pool(nworkers)
	<-done

	finish := time.Now()
	diff := finish.Sub(start)
	fmt.Printf("total: %.2v sec\n", diff)
}

type Job struct {
	id   int
	data int
}

type Result struct {
	job  Job
	data int
}

func allocate(njobs int) {
	for i := 0; i < njobs; i++ {
		job := Job{id: i, data: 2}
		jobs <- job
	}
	close(jobs)
}

func pool(nworkers int) {
	var wg sync.WaitGroup
	for i := 0; i < nworkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		out := Result{job, sqrt(job.data)}
		results <- out
	}
	wg.Done()
}

func sqrt(x int) int {
	time.Sleep(2 * time.Second)
	return x * x
}

func result(done chan bool) {
	for r := range results {
		j := r.job
		fmt.Printf("job(%v, %v) = %v\n", j.id, j.data, r.data)
	}
	done <- true
}
