FROM golang:1.18.3


WORKDIR $GOPATH/src/github.com/conglindeng/integration

COPY . .

RUN go get 