package main

import (
	"fmt"
)

func Tap(f func(int)) *tapFilter {
	return &tapFilter{
		f:   f,
		out: make(chan int),
	}
}

type tapFilter struct {
	f   func(x int)
	in  chan int
	out chan int
}

func (t *tapFilter) SetIn(in chan int) {
	t.in = in
}

func (t *tapFilter) Out() chan int {
	return t.out
}

func (t *tapFilter) Run() {
	for x := range t.in {
		t.f(x)
		t.out <- x
	}
	close(t.out)
}

func Func(f func(int) int) *funcFilter {
	return &funcFilter{
		f:   f,
		out: make(chan int),
	}
}

type funcFilter struct {
	f   func(int) int
	in  chan int
	out chan int
}

func (ff *funcFilter) SetIn(in chan int) {
	ff.in = in
}

func (ff *funcFilter) Out() chan int {
	return ff.out
}

func (ff *funcFilter) Run() {
	for x := range ff.in {
		ff.out <- ff.f(x)
	}
	close(ff.out)
}

func Square() *sqr {
	return &sqr{
		out: make(chan int),
	}
}

type sqr struct {
	in  chan int
	out chan int
}

func (s *sqr) SetIn(in chan int) {
	s.in = in
}

func (s *sqr) Out() chan int {
	return s.out
}

func (s *sqr) Run() {
	for x := range s.in {
		s.out <- x * x
	}
	close(s.out)
}

func Range(from, to int) *rng {
	return &rng{
		from: from,
		to:   to,
		out:  make(chan int),
	}
}

type rng struct {
	from, to int
	out      chan int
}

func (r *rng) SetIn(_ chan int) {}

func (r *rng) Run() {
	for i := r.from; i < r.to; i++ {
		r.out <- i
	}
	close(r.out)
}

func (r *rng) Out() chan int {
	return r.out
}

type Filter interface {
	SetIn(chan int)
	Out() chan int
	Run()
}

func Pipe(filters ...Filter) *pipe {
	return &pipe{filters: filters}
}

type pipe struct {
	filters []Filter
}

func (p *pipe) Run() {
	r := p.filters[0]
	for i := 1; i < len(p.filters); i++ {
		f := p.filters[i]
		f.SetIn(r.Out())
		r = f
	}
	for _, f := range p.filters {
		go f.Run()
	}
}

func (p *pipe) Out() chan int {
	return p.filters[len(p.filters)-1].Out()
}

func main() {

	p := Pipe(
		Range(0, 10),
		Square(),
		Func(func(x int) int { return x + 1 }),
		// Tap(func(x int) { fmt.Println(x) }),
	)
	p.Run()

	for x := range p.Out() {
		fmt.Print(x, " ")
	}
	fmt.Println()
}
