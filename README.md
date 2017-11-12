# docker-compose-generator-cli
This is very initial version of simple CLI tool written in golang

My dream is to allow simple commands like:

dcgc tool redis

generating/appending following docker-compose config:

service redis:

    image: redis
    ports:
       - 6379:6379
    volumes:
       - redis:/data
       