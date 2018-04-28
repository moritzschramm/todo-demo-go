# Todo Demo application

Simple Todo web application written in Go and Vue.js.
Preview: (Front-end only): [moritz.schramm.com/todo-vuejs](https://moritz-schramm.com/todo-vuejs)

## Requirements
 - `docker` 
 - `docker-compose`

## Usage
```
docker-compose up
```
This will start 2 docker containers (one for the server, the other one for the database). The first execution will take a bit longer since the mysql container has to setup the database. All data will be stored in `./data` (and doesn't get lost when you stop docker-compose)