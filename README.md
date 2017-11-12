# docker-compose-generator-cli
This is very initial version of simple CLI tool written in golang.
Main goal is to provide simple automation tool for prototyping microservices/tools.
It's also learning project to improve my skills in golang.
My dream is to allow simple commands like:

     dcgc tool redis
     
generating/appending following docker-compose config:

    redis:
        image: redis
        ports:
           - 6379:6379
        volumes:
           - redis_data:/data
       
       
I'm gonna describe more features in future.


    
