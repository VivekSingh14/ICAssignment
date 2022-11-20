# ICAssignment
Short Url builder

# How to run it on local: 
    use the following steps and hit the given url on browser or on postman because Method type is GET only 
go run -tags dynamic cmd/main.go


Endpoint:  localhost:8080/short/v1/create/{url}
    where url = whatever you want to convert into short url
Method: GET

# Steps to run it using docker

1. pull image using below command
        #docker pull vivek1411/shorturl-latest

2. run the container using below command
      #docker run -d -p 8080:8080 --name <name> vivek1411/shorturl-latest

3. hit it using curl or browser or postman.

