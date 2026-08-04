FROM golang:1.22-alpine

WORKDIR /app

RUN apk add --no-cache git gcc musl-dev ca-certificates

COPY go.mod go.sum* ./

RUN go mod download 2>/dev/null || true

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o app ./cmd/api

EXPOSE 8080

CMD ["./app"]
