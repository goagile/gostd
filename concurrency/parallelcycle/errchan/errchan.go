package main

import (
	"sync"
	"log"
	"fmt"
)

func main() {
	
	errors := make(chan error)

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			errors <- calculate(i)
		}(i)
	}

	go func() {
		wg.Wait()
		close(errors)
	}()

	for err := range errors {
		if err != nil {
			log.Println(err)
		}
	}

}

func calculate(i int) error {
	if i > 2 {
		return fmt.Errorf("err %v", i)
	}
	return nil
}