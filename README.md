# URL Shortener

URL Shortener is an api to allow user to convert their full-length URL into a short URL, convert back into their original URL also redirect to the Original URL.
This API projects comes with environment variable to make it easier to deploy everywhere.

# Requirement
- Golang >= 1.12
- Mongodb
- Docker

# Environtment Variable
- appport (Port for running api)
- dbhost (Database Host)
- dbport (Database Port)
- db (Database Name)
- dbuser (Database Username) *#Optional*
- dbpass (Database Password) *#Optional*

# API
There are 3 APIs to access:
```
/register?{long url}
/retrieve?{short url}
/redirect?{short url}
```
#### Register
To register long URL and return json data with details of short URL
```
/register?{long url}
```
example:
![register](https://github.com/kotopanjang/url_shortener/blob/master/resources/register%201.png)

#### Retrieve
Access short URL and return json data with details of oiginal/long URL
```
/retrieve?{short url}
```
example:
![retrieve](https://github.com/kotopanjang/url_shortener/blob/master/resources/retrieve%201.png)


#### Redirect
Access short URL and redirrect to original/long URL
```
/redirect?{short url}
```
example:
![redirect 1](https://github.com/kotopanjang/url_shortener/blob/master/resources/redirect%201.png)


When you hit the url


![redirect 2](https://github.com/kotopanjang/url_shortener/blob/master/resources/redirect%20result%201.png)

# How It Works?
#### Register
When someone register the long URL, it will check to database wether the long URL is already registered on the database or not. And also checking the expired time. As default it will expired in 1 day.
If the long URL is found and not expired, it will take that data and return it.
If the long URL is not found or expired, it will insert new data with expired 1 day.
To avoid the duplicate active short URL stored in database, it will prepare 25 Short URL and do the checking on the database.

Result
```
{
  data: "localhost:4444/retrieve?enmjv4Ly",
  error: false,
  message: null
}
```

#### Retrieve
It willl check on the database wether the short URL is expired or not. If expired, it will return json invalild URL. And if the data is not expired then it will show the full URL in data section.

Result Success
```
{
  data: "https://www.google.com/search?q=aqilliz&oq=aqill&aqs=chrome.0.69i59j46j69i57j0l2j69i60l3.1720j0j7&sourceid=chrome&ie=UTF-8",
  error: false,
  message: null
}
```
Result Invalid URL
```
{
  data: "",
  error: true,
  message: "Invalid url"
}
```

#### Redirect
The flow between `retrieve` and `redirect` is the same. But in `redirect`, it will redirect to full URL directy. If the URL is invallid, it will return invalid URL json

Result Invalid URL
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
