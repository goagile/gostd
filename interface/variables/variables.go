package main

import "fmt"

type Z struct {
	I int
	B bool
}

const format = "%4v %20p %15T  %v\n"

func main() {
	var i int = 25
	fmt.Printf(format, "i", &i, i, i)

	var fi interface{} = 3.4
	fmt.Printf(format, "fi", &fi, fi, fi)

	var n interface{} = nil
	fmt.Printf(format, "n", &n, n, n)

	var b bool = true
	fmt.Printf(format, "b", &b, b, b)

	var s []int = []int{1, 2, 3}
	fmt.Printf(format, "s", &s, s, s)

	var z Z = Z{23, true}
	fmt.Printf(format, "z", &z, z, z)

	var pz *Z = &Z{23, true}
	fmt.Printf(format, "pz", &pz, pz, pz)

	var spz []*Z = []*Z{&z, pz}
	fmt.Printf(format, "spz", &spz, spz, spz)

	var si []interface{} = []interface{}{i, fi, n, b, s, z, pz, spz}
	fmt.Printf(format, "si", &si, si, si)
}
