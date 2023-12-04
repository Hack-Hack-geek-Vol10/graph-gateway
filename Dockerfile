FROM golang:1.21.4 AS builder
WORKDIR /app

COPY . .

RUN go build -o graphql-server ./cmd/app/main.go

FROM ubuntu:latest AS runner

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder ./app/graphql-server /app
COPY --from=builder ./app/*.json /app

ENV SERVER_ADDR=8080
ENV USER_SERVICE_ADDR="localhost:8081"
ENV PROJECT_SERVICE_ADDR="localhost:8082"
ENV MEMBER_SERVICE_ADDR="localhost:8083"
ENV IMAGE_SERVICE_ADDR="localhost:8084"
ENV TOKEN_SERVICE_ADDR="localhost:8085"
ENV NEW_RELIC_LICENSE_KEY=
ENV NEW_RELIC_APP_NAME=
ENV NEW_RELIC_CODE_LEVEL_METRICS_ENABLED=

EXPOSE ${SERVER_ADDR}

CMD ["./graphql-server"]