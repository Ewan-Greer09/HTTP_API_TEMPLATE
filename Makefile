#runs main.go and attaches nodemon
start-jobboard:
	./start.cmd

start-validationserver:
	./start-validationserver.cmd

start-authserver:
	./start-authserver.cmd

start-all: start-jobboard start-validationserver start-authserver

build:
	go build

clean:
	go clean
