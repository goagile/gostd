package main

type empty struct{}

type semaphore chan empty

func (s semaphore) acquare(n int) {
	e := empty{}
	for i := 0; i < n; i++ {
		s <- e
	}
}

func (s semaphore) release(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

func (s semaphore) Lock() {
	s.acquare(1)
}

func (s semaphore) Unlock() {
	s.release(1)
}

func (s semaphore) Wait(n int) {
	s.acquare(n)
}

func main() {

}
