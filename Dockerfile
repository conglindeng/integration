FROM golang:1.18.3

WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

ENV MODIFY=2022-12-27

COPY src/ ./src/

RUN cd /app/src/github.com/conglindeng/integration && go mod download  && go build -o integration

EXPOSE 9090

CMD [ "/app/src/github.com/conglindeng/integration/integration" ]