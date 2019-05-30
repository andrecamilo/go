# para funcionar, deve-se rodar os seguintes comandos

go get github.com/crgimenes/goconfig


export DOMAIN=$DOMAIN:teste_domain
export MONGODB_HOST=$MONGODB_HOST:teste_mongodb_host



export DOMAIN="teste_domain"
export MONGODB_HOST="teste_mongodb_host"




//docker run --name config --restart=always -p 50002:80 -v trip_planner:/app/Inlog.Trip.Planner.WebApi -e ASPNETCORE_ENVIRONMENT=Staging -d 



docker build -t config:1 .
docker run config:1