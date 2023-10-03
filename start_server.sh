#!/bin/bash
export GIN_MODE=release 
# Run swag init to generate Swagger documentation
swag init .
# Start your Go Gin API server
go run main.go
