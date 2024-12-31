# How to scale websockets

## To run locally
```
# start a redis instance in local using docker or directly install redis
docker run -p 6379:6379 --name redis-local redis

# run the golang main file
go run cmd/chat/main.go 

# make sure to change the image id and tag
# to create a docker image
make image

# to push image
make push_image

# to create a docker image for arm processor
make image_arm

# to push image
make push_image_arm

# to create a docker image for amd process
make image_amd

# to push image
make push_image_amd

# use docker compose to test the app with haproxy
docker-compose up


# to add additional server add another block in the docker-compose.yaml and haproxy.conf

# docker-compose.yaml
  chat-worker-1:
    image: abhishek1009/group-chat:latest
    container_name: chat-worker-1
    environment:
      - SERVER_NAME=chat-worker-1
      - SERVER_ADDRESS=0.0.0.0
      - REDIS_HOST=redis
      - REDIS_PORT=6379
    networks:
      - chat-network
    depends_on:
      - redis

# haproxy.conf
backend all
    server s1 chat-worker-1:8080
```

## deploy locally using minikube with ingress
```
# to enable ingress in minikube 
minikube addons enable ingress

# update /etc/hosts file and entry for the ingress host
sudo echo "127.0.0.1	chat.local" >> /etc/hosts

## run the kubernetes manifest file
kubectl apply -f deployment/chat-worker-manifest.yaml
```


## Live reloading
```
# visit air repository by the following link:
https://github.com/cosmtrek/air

# install the air tool with go:
go install github.com/cosmtrek/air@latest

# initialize project:
air init

# start project with existing toml file
air -c .air.toml
```

## Resources
- [Scaling Websockets with Redis, HAProxy and Node JS - High-availability Group Chat Application](https://www.youtube.com/watch?v=gzIcGhJC8hA)
  - [Source code](https://github.com/hnasr/javascript_playground/tree/master/ws-live-chat-system)
- [Piyush Garg](#)
  - [Build Scaleable Realtime Chat App with NextJS and NodeJS Tutorial](https://www.youtube.com/watch?v=CQQc8QyIGl0)
  - [Build Scaleable Realtime Chat App with Kafka and Postgresql](https://www.youtube.com/watch?v=Rat7ORbBDN8)