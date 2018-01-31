
Начало работы с go

	$ cd ~
	$ mkdir go
	$ cd go
	$ mkdir src
	$ mkdir bin
	$ mkdir pkg
	$ export GOPATH=$HOME/go/
	$ export GOBIN=$GOPATH/bin
	$ cd src
	$ mkdir hello
	$ cd hello
	$ touch hello.go
	$ subl hello.go
	$ go install
	$ export PATH=$GOBIN:$PATH
	$ $GOBIN/hello
	$ hello

Создание модулей go

	$ mkdir string
	$ cd string
	$ touch string.go
	$ subl string.go
	$ go build
	$ go install
	$ cd ..
	$ cd hello/
	$ subl hello.go
	$ cd ..
	$ go install

Установка модулей

	$ go get github.com/julienschmidt/httprouter
