docker network create -d bridge temprabbitnet
docker run -d --name temprabbit --restart always -p 5672:5672 -p 0.0.0.0:15672:15672 rabbitmq:3-management
#docker network connect temprabbitnet temprabbit