package main

import (
	"regexp"
	"strings"
	"testing"
)

const s = "aaa bbb Hello ccc ddd Hello ggg"

var re = regexp.MustCompile("Hello")

func BenchmarkNonCompile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		regexp.MustCompile("Hello").FindAllString(s, -1)
	}
}

func BenchmarkPreCompile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		re.FindAllString(s, -1)
	}
}

func BenchmarkSubString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Contains(s, "Hello")
	}
}
