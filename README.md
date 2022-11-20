# ICAssignment
Short Url builder

# How to run it on local: use the following steps and hit the given url on browser or on postman because Method type is GET only 
go run -tags dynamic cmd/main.go


Endpoint:  localhost:8080/short/v1/create/{url}
    where url = whatever you want to convert into short url
Method: GET

# How to run it using docker, will be adding the dockerfile as well.
