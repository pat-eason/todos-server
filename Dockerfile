FROM golang:1.21.4-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN GOARCH=arm64 GOOS=linux go build -o bin/todo-server-linux cmd/api.go

EXPOSE 8080

CMD ["todo-server-linux"]