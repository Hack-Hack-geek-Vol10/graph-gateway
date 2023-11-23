generate:
	gqlgen generate

run:
	go run cmd/app/main.go

.PHONY: generate run