
FROM golang:1.13-alpine

RUN apk update && apk add git bash make vim

COPY . /app

WORKDIR /app

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build

CMD ["./main"]

EXPOSE 8080
