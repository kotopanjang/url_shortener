# URL Shortener

URL Shortener is an api to allow user to convert their full-length url into a short link. And also convert back into their original URL.

# Requirement
- Golang >= 1.12
- Mongodb
- Docker (Optional)

# Environtment Variable
- appport (Port for running api)
- dbhost (Database Host)
- dbport (Database Port)
- db (Database Name)
- dbuser (Database Username) *#Optional*
- dbpass (Database Password) *#Optional*

# API
There are 3 APIs to access


# Installation
Get all the source code with go get.
> go get github.com/kotopanjang/url_shortener


## Running without Docker File
Before you run the api, make sure your mongodb is already set up. If not, you can open terminal and use this command.
```
mongod --dbpath /data/db
```

After setup the database, now its time to setup the api. First you need to go to "url_shortener" folder and build the project with this command.
```
go build -o url_shortener
```

Then you can run the api with environtment variable above.
```
./url_shortener appport=2020 dbhost=localhost dbport=27017 db=url_shortener 
```

## Running with Docker File
Inside "url_shortener" folder you will find file **Dockerfile** that already configured. But before you run the api, you need to run mongodb in your docker container with the same network.

First you need to create docker network
```
docker network create my_network
```

Now get mongodb docker image
```
docker pull mongo:latest
```

After you get the mongodb docker image, now you can run it on the container and expose the port. 
Also run under the same network  
```
docker run --name mongo-docker -p 27017:27017 --network my_network mongo:latest
```

Run the api
```
docker run -it -e appport=2020 -e dbhost=mongo-docker -e dbport=27017 -e db=testing_aqilliz --network my_network -p 2020:2020 url_shortener
```

## Running with Docker Compose
On the root folder, you will find docker-compose.yml that already configured.
Build docker compose
```
docker-compose build
```
Run docker compose
```
docker-compose up
```
