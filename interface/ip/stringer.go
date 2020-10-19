package main

import (
	"fmt"
	"strings"
)

func main() {
	addrs := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"google dns": {8, 8, 8, 8},
	}

	for name, ip := range addrs {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	s := make([]string, 4)
	for i, v := range ip {
		s[i] = fmt.Sprintf("%v", v)
	}
	return strings.Join(s, ".")
}
