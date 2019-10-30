FROM golang:1.13-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN GOOS=linux GOARCH=amd64 go build -o main .

CMD ["/app/main"]