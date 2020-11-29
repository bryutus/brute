FROM golang:latest

RUN mkdir /app
WORKDIR /app

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y vim

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
CMD ./bin/air

RUN go get github.com/gin-gonic/gin && \
    go get github.com/go-sql-driver/mysql && \
    go get -u golang.org/x/sync/errgroup && \
    go get github.com/jinzhu/gorm
