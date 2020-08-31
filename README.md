# URL Shortener

URL Shortener is an api to allow user to convert their full-length URl into a short link. And also convert back into their original URL.
It comes with environment variable to make it easier to deploy everywhere.

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
There are 3 APIs to access `register`, `retrieve`, `redirect`.

#### Register
To register long URl and return json data with details of short URl
```
/register?{long url}
```
example:
![register](https://github.com/kotopanjang/url_shortener/blob/master/resources/register%201.png)

#### Retrieve
Acecss your short URl and return json data with details of oiginal/long URl
```
/retrieve?{short url}
```
example:
![retrieve](https://github.com/kotopanjang/url_shortener/blob/master/resources/retrieve%201.png)


#### Redirect
Acecss your short URl and redirrect to original/long URl
```
/redirect?{short url}
```
example:
![redirect 1](https://github.com/kotopanjang/url_shortener/blob/master/resources/redirect%201.png)

`when you hit the url`

![redirect 2](https://github.com/kotopanjang/url_shortener/blob/master/resources/redirect%20result%201.png)

# How It Works?
#### Register
When someone register the long URL, it will check to database wether the long URL is already registered on the database or not. And also checking the expired time. As default it will expired in 1 day.
If the long URL is found and not expired, it will take that data and return it.
If the long URl is not found or expired, it will insert new data with expired 1 day.
To avoid the duplicate random string, it will prepare 25 random string and do the checking on the database.
```
when the api
```

#### Retrieve
It willl check on the database wether the short URl is expired or not. If expired, it will return json invalild URl. And if the data is not expired then it will show the full URl in data section.
`Invalid URL`
```
{
  data: "",
  error: true,
  message: "Invalid url"
}
```

#### Redirect
The flow between `retrieve` and `redirect` is the same. But in `redirect`, it will redirect to full URL directy. If the URL is invallid, it will return invalid url json
`Invalid URL`
```
{
  data: "",
  error: true,
  message: "Invalid url"
}
```

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

Dockerize the api under the same network
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
