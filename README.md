# Lets-Go

Setting up an HTTP server with Golang.

Building a simple server with **"net/http"** library in Golang.<br />
This is a simpe server with two routes, the main idea behind this project<br />
was to learn the basics of HTTP in Golang.

## Tools
- go (version 1.17)
- k6

## Run the project
Clone the project, then run the following command:
```shell
go run cmd/server/main.go
```

And you will get this response:
```shell 
Server is listening on 127.0.0.1:8080 ...
```

Now you can test the server with **k6** with the following command:
```shell
k6 run ./api/k6/script.js
```

Or using Goland HTTP Client with the following command:
```shell 
go run cmd/client/main.go
```