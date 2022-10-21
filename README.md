## GO URL Shortner
 GO URL Shortner is a  URL shortener service that will accept a URL as an argument over a REST API and return a shortened URL as a result



## Requirements

* **Go:** `1.19`
* **Gin:** `1.8.1`
* **Docker:**  (optional)


## Installation and Running the service

#### Option 1:  Cloning from git hub and installing on systems with  requirements mentioned above.

``` shell
$git clone git@github.com:aicelerity-golang/go-url-shortner.git
$cd go-url-shortner
$go mod tidy
$go mod verify
$go run .
```
The application GO URL Shortner will start and run on port 5000. You will see below log if app is running successful.

``` shell
[GIN-debug] Listening and serving HTTP on :5000

```
#### Option 2: Using Docker

### Getting the Short Url Service

* REST API `/url-shortner ` running on localhost port 5000.

#### Option 1. Using curl command
Given below an example of getting the Short URL using curl command.

``` shell
$ curl --request POST --data '{"long-url": "https://blogs.blackberry.com/en/2022/10/bianlian-ransomware-encrypts-files-in-the-blink-of-an-eye"}'  http://localhost:5000/url-shortner

```

If it is successfull, you should get below output.


``` shell
{"message":"Here comes the Short Url...","short-url":"http://localhost:5000/RX9wbW9WtQ"}

```
#### Option 2. Using Postman

Open Post man App.
Selct New
Select  HTTP Request
Under Request select POST from drop down
Enter the URL (http://localhost:5000/url-shortner)
Select Body then raw amd add the json ( example: {"long-url": "https://blogs.blackberry.com/en/2022/10/bianlian-ransomware-encrypts-files-in-the-blink-of-an-eye"
}) 
Click Send and you should get the Short URL as response.


### Source Files


### Tests
