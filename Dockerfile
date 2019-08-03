FROM golang:1.12.7

WORKDIR /go/src/app
COPY . .

RUN go build -o main

CMD ["./main"]