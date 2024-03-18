FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go build ./cmd/main/main.go

CMD ["./main"]
