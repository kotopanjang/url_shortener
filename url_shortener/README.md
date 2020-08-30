#URL Shortener

URL Shortener is an api to allow user to convert their full-length url into a short link. And also convert back into their original URL.

#Requirement
- Golang >= 1.12
- Mongodb

#Environtment Variable
- appport (Port for running api)
- dbhost (Database Host)
- dbport (Database Port)
- db (Database Name)
- dbuser (Database Username) *#Optional*
- dbpass (Database Password) *#Optional*

#Installation
Get all the source code with go get.
> go get github.com/kotopanjang/url_shortener


##Running without Docker File
Before you run the api, make sure your mongodb is already set up. If not, you can open terminal and use this command.
> mongod --dbpath /data/db

After setup the database, now its time to setup the api. First you need to build the api with this command.
> go build -o url_shortener

Then you can run the api with environtment variable above.
> ./url_shortener appport=2020 dbhost=localhost dbport=27017 db=url_shortener 

##Running with Docker File
Inside "url_shortener" folder  you will find **Dockerfile** that already configured. But before you run the api, you need to run the mongodb in your docker first.
>


####Running with Docker Container
