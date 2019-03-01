docker run -d --hostname my-rabbit --name some-rabbit rabbitmq:3
docker logs some-rabbit
docker run -d --hostname my-rabbit --name some-rabbit rabbitmq:3-management







docker run -d — hostname my-rabbit — name rabbit13 -p 8080:15672 -p 5672:5672 -p 25676:25676 rabbitmq:3-management