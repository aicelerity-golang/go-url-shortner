## GO URL Shortner
GO URL Shortner is a  URL shortener service that will accept a URL as an argument over a REST API and return a shortened URL as a result.


## Requirements

* **Go:** `1.19`
* **Gin:** `1.8.1`
* **Docker:**  (optional)


## Installation and Running the  URL shortener service

#### Option 1:  Cloning from git hub and installing on systems with  requirements mentioned above.

``` shell
$git clone git@github.com:aicelerity-golang/go-url-shortner.git
$cd go-url-shortner
$go mod tidy
$go mod verify
$go run .
```
Or we could use go build instead of go run like below.
``` shell
$go build .
$./go-url-shortner
```
The application GO URL Shortner will start and run on port 5000. We can see below log once the app is running successful.

``` shell
[GIN-debug] Listening and serving HTTP on :5000

```
#### Option 2: Using Docker 
Docker can be download from [here](https://docs.docker.com/get-started/#download-and-install-docker).
Docker or Docker Desktop need to be installed as a prerequisite for below steps.

Follow below steps to get the application running as a container.

``` shell
docker pull rajeshmathewk/go-url-router:latest
docker run -d -p 5000:5000 -it --name go-url-router rajeshmathewk/go-url-router
docker ps
docker logs -f  <container ID>
```
We can see below log once the app is running in container.
``` shell
[GIN-debug] Listening and serving HTTP on :5000

```

### Requesting the URL shortener Service

* REST API `/url-shortner ` running on localhost port 5000.

#### Option 1. Using curl command
Given below an example of getting the Short URL using curl command.

``` shell
$ curl --request POST --data '{"long-url": "https://blogs.blackberry.com/en/2022/10/bianlian-ransomware-encrypts-files-in-the-blink-of-an-eye"}'  http://localhost:5000/url-shortner

```
If the URL Shortner service is running we should get below output.

``` shell
{"message":"Here comes the Short Url...","short-url":"http://localhost:5000/RX9wbW9WtQ"}

```
#### Option 2. Using Postman

Open Postman App. Navigate to  New ->  HTTP Request. Under Request select POST from drop down . Enter the URL: http://localhost:5000/url-shortner.

Select Body -> Then raw and add the json data  example: 
{"long-url": "https://blogs.blackberry.com/en/2022/10/bianlian-ransomware-encrypts-files-in-the-blink-of-an-eye"} . Click Send and we should get the Short URL as json response shown below.

``` shell
{
    "message": "Here comes the Short Url...",
    "short-url": "http://localhost:5000/RX9wbW9WtQ"
}
```


### Source Files

*  `main.go`           : will handle the HTTP requests using Gin Framework

* ***utils*** 
  - `urlshortner.go`   : does the shortnening of given Long URL; hash and encode to a short URL.
  
* ***handler***  
  - `filehandler.go`   : will handle the File operations; reading and writing of URLs in File.
  - `reqhandler.go`    : will handle the API Request Service; request and response of URL Service.
  - `cachehandler.go`  : does the caching of URLs; reading and writing of URLs in memory.


### Tests 
*  `handler/filehandler_test.go`  : runs the tests for File handler Operations.
*  `utils/urlshortner_test.go`    : runs the tests for URL shortner ; genrerating a Short URL.
