FROM golang:alpine

RUN mkdir -p /go/src && \
    cd /go/src

WORKDIR /go/src/
COPY src .

RUN apk upgrade --update &&\
    apk add vim curl

CMD ["go". "run", "main.go"]
