FROM golang:1.25-alpine

WORKDIR /go/src/app

COPY . .

RUN go build -o main main.go

CMD ["./main"]
