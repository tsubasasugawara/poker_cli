FROM golang:latest

RUN mkdir -p /go/src && \
    cd /go/src

ENV ROOT=/go/src
WORKDIR ${ROOT}

COPY src .

RUN apt update && apt upgrade -y &&\
    apt install vim curl -y

RUN go get github.com/pkg/term/termios &&\
    go get github.com/nsf/termbox-go &&\
    go get github.com/gorilla/websocket@latest &&\
    go mod tidy

RUN echo "umask 000" >> ~/.profile
