package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {

	ncpu := runtime.NumCPU()
	fmt.Println("CPU:", ncpu)
	runtime.GOMAXPROCS(ncpu)

	names := []string{
		"Петров Иван", 
		"Сидоров Степан", 
		"Антонов Егор",
		"Иванов Антон",
		"Степанов Петр",
		"Сафронов Алексей",
	}

	jobs := make(chan Job)
	
	results := make(chan Result)

	done := make(chan struct{}, ncpu)

	//
	// allocate jobs
	//
	go func(jobs chan Job, names []string) {
		for id, name := range names {
			jobs <- Job{id, name}
		}
		close(jobs)
	}(jobs, names)

	//
	// do jobs
	//
	for w := 0; w < ncpu; w++ {
		go func(w int, jobs chan Job, results chan Result, done chan struct{}) {
			for j := range jobs {
				r := Result{}
				r.JobID = j.ID
				r.WorkerID = w
				parts := strings.Split(j.Name, " ")
				r.First = parts[0]
				r.Last = parts[1]
				results <- r
			}
			done <- struct{}{}
		}(w, jobs, results, done)
	}

	//
	// awaiting for results
	//
	go func() {
		for w := 0; w < ncpu; w++ {
			<-done
		}
		close(results)
	}()
	
	//
	// read results
	//
	for r := range results {
		fmt.Printf("%#v\n", r)
	}
	fmt.Println()

}

type Job struct {
	ID int
	Name string
}

type Result struct {
	JobID int
	First, Last string
	WorkerID int
}
