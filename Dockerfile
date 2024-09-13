FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .env .

RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]
