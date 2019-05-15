FROM golang:1.12

WORKDIR /go-docker
COPY . /go-docker

# RUN go get -d -v ./...
# RUN go install -v ./...

RUN go get "github.com/gorilla/mux"
RUN go get "gopkg.in/natefinch/lumberjack.v2"

EXPOSE 8080
# CMD ["docker"]


# FROM golang
# COPY . /hello


ENTRYPOINT ["/go-docker"]