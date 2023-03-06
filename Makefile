#runs main.go and attaches nodemon
all: run-jobboard run-validationserver run-authserver

run-jobboard:
	nodemon --exec go run /services/jobboard/main.go

run-validationserver:
	nodemon --exec go run ./services/validationserver/main.go
