FROM golang:1.17.0 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go get && go build -o crawler main.go

FROM debian:buster-slim AS production
WORKDIR /app
COPY --from=builder /app/crawler /app/crawler
ENTRYPOINT ["./crawler"]