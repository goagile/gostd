package main

import "fmt"

type Weekday int

const (
  Sunday Weekday = iota
  Monday
  Tuesday
)

const (
  pow2_0 int = 2 * iota
  pow2_1
  pow2_2
  pow2_3
  pow2_4
)

func main() {
  day := Monday
  fmt.Printf("day = %v\n", day)

  p2_2 := pow2_2
  fmt.Printf("p2_2 = %v\n", p2_2)
}
