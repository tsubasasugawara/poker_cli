FROM golang:alpine

RUN mkdir -p /go/src && \
    cd /go/src

WORKDIR /go/src
COPY ./src /go/src/

RUN apk upgrade --update &&\
    apk add vim curl gcc musl-dev

RUN go get github.com/gin-gonic/gin@latest &&\
    go install github.com/cosmtrek/air@latest &&\
    go mod tidy

EXPOSE ${WEBSITES_PORT}
