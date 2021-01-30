# !/bin/bash

echo "--> run go benchmark"
go test -bench . -cpuprofile=cpu.out

echo "--> run go pprofile"
go tool pprof fact.test cpu.out

