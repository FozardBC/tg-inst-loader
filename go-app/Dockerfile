FROM golang:1.22-alpine as builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o tg-bot ./cmd


CMD ["./tg-bot"]