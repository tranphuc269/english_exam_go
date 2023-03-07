FROM golang:1.19-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app
COPY go.mod go.sum ./


RUN go mod download

#RUN go build -o main server.go

ENTRYPOINT ["go","run","./server.go"]

