@echo off
nodemon -e go -x "go run .\jobboard\main.go"
nodemon -e go -x "go run .\authserver\main.go"
nodemon -e go -x "go run .\validationserver\main.go"