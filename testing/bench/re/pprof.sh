# !/bin/bash

# run benchmark with profile of CPU to file "cpu.out"
go test -bench . -benchmem -cpuprofile=cpu.out re_test.go

# run profile from file
go tool pprof main.test cpu.out
