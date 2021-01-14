package main


import (
	"fmt"
	"math"
)


func main() {
	fmt.Println(Sqrt(-9))
	fmt.Println(Sqrt(9))
}


func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNeg(x)
	}
	return math.Sqrt(x), nil
}

type ErrNeg float64

func (e ErrNeg) Error() string {
	return fmt.Sprintf("cannot be negative: %v", float64(e))
}
