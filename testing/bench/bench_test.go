package main

import (
	"testing"
)

var examples = []struct{ result, n int }{
	{1, 0},
	{1, 1},
	{2, 2},
	{6, 3},
	{24, 4},
	{120, 5},
}

func factRecur(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return n * factRecur(n-1)
}

func TestFactRecur(t *testing.T) {
	for _, e := range examples {
		if e.result != factRecur(e.n) {
			t.Fatalf("error with n: %v\n", e.n)
		}
	}
}

func BenchmarkFactRecur(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factRecur(5)
	}
}

func factCicle(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	p := 1
	for i := 2; i <= n; i++ {
		p *= i
	}
	return p
}

func TestFactCicle(t *testing.T) {
	for _, e := range examples {
		if e.result != factCicle(e.n) {
			t.Fatalf("error with n: %v\n", e.n)
		}
	}
}

func BenchmarkFactCicle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factCicle(5)
	}
}
