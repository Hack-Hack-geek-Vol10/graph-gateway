generate:
	gqlgen generate

run:
	go run server.go

.PHONY: generate run