FROM golang:1.21.4 AS builder
WORKDIR /app

COPY . .

RUN go build -o graphql-server ./cmd/app/main.go

FROM alpine:3.14.2 AS runner

WORKDIR /app

COPY --from=builder ./app/graphql-server /app

ENV SERVER_ADDR=8080
ENV USER_SERVICE_ADDR="localhost:8081"
ENV PROJECT_SERVICE_ADDR="localhost:8082"
ENV MEMBER_SERVICE_ADDR="localhost:8083"
ENV IMAGE_SERVICE_ADDR="localhost:8084"
ENV TOKEN_SERVICE_ADDR="localhost:8085"

EXPOSE ${SERVER_ADDR}

CMD ["./graphql-server"]