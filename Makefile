build:
	go build -o bin/xdostr xdostr.go

run:
	go run xdostr.go

install:
	go build -o /usr/local/bin/xdostr xdostr.go
