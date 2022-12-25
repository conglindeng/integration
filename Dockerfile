FROM golang:1.18.3


WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# COPY src/*

RUN go get 