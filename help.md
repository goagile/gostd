
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
