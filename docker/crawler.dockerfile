FROM golang:1.17.0 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go get && go build -o google-crawler .

FROM debian:buster-slim AS production
WORKDIR /app
COPY --from=builder /app/google-crawler /app/google-crawler
RUN chmod +x ./google-crawler
CMD ["./google-crawler"]