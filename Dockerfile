FROM golang:1.14

WORKDIR /app

ADD ./go.mod ./go.sum ./

CMD go mod download