FROM golang:1.22-alpine

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum* ./

RUN go mod download 2>/dev/null || true

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api

EXPOSE 8080

CMD ["./app"]
