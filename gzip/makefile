

go-pack: 
	go run ./pack/pack.go ./pack/*.txt
	cp ./pack/*.gz ./unpack

go-unpack:
	go run ./unpack/unpack.go ./unpack/*.gz

clear:
	rm ./pack/*.gz
	rm ./unpack/*.gz
	rm ./unpack/*.txt
