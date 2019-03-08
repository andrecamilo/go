FROM golang:1.6.2

COPY . /server

RUN go get github.com/nats-io/nats
RUN go build server.go

EXPOSE 8080
ENTRYPOINT ["/server/server"]