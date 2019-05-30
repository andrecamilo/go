
export GO111MODULE=on
go mod init
go build

docker build -t go-docker .
docker run -d -p 8080:8080 go-docker

curl http://localhost:8080?name=Andre
Hello, Andre