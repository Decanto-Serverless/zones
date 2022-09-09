SHELL=/bin/sh

buildWin:
	GOOS=windows GOARCH=amd64 go build -o handler main.go

build:
	go build -o handler main.go

run: build
	func start

deploy: buildWin
	func azure functionapp publish decanto-zones