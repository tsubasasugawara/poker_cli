FROM golang:latest

RUN mkdir -p /go/src && \
    cd /go/src

ENV ROOT=/go/src
WORKDIR ${ROOT}

COPY src .

RUN apt update && apt upgrade -y &&\
    apt install vim curl -y

RUN echo "umask 000" >> ~/.profile
