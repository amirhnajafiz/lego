<p align="center">
    <img src=".github/assets/golang.jpeg" alt="logo" />
</p>

<h1 align="center">
Lets Go
</h1>

This is my first project with Go programming language. In this project I set up an HTTP server with Golang and then
test that server with two different ways. In the first case we are using a Golang client, in the second case we are using Grafana k6 testing.

This project is perfect for people who want to start learning Golang.

## What did I learn?
- golang files and package structure
- net/http package
- k6 testing

## Project structure in Golang 
Every Golang project has a **go.mod** file. This file manages the modules that we import and use 
in our Golang projects. Remember that module name should be like your github repository address.
```
module github.com/USERNAME/REPOSITORY
```

Now we have a _CMD_ which contains all of our Golang application codes. In this folder we have an _http_ module which manages 
the server handlers and server responses. The other module is _server_ itself that manages the application.
```
|_ client/
    |_ main.go
|_ cmd/
    |_ http/
        |_ handler/
        |_ response/
    |_ server/
        |_ main.go
```

## Run the project on your local
Clone the project, then run the following command:
```shell
go run cmd/server/main.go
```

And you will get this response:
```shell 
Server is listening on 127.0.0.1:8080 ...
```

This command creates a **HTTP server** for you. Now you can access this server on _localhost:8080_.

## Connect to our server
### Grafana k6
#### What is k6?
Grafana k6 is an open-source load testing tool that makes performance testing easy and productive for engineering teams. k6 is free, developer-centric, and extensible.

Using k6, you can test the reliability and performance of your systems and catch performance regressions and problems earlier. k6 will help you to build resilient and performant applications that scale.

k6 is developed by Grafana Labs and the community.

Now you can test the server with **k6** using the following command:
```shell
k6 run api/k6/script.js
```

And results will be something like this:
```shell
running (0m23.7s), 00/20 VUs, 46 complete and 8 interrupted iterations
default ✗ [====>---------------------------------] 08/20 VUs  0m23.7s/3m00.0s

     ✓ status is 200

     checks.........................: 100.00% ✓ 97       ✗ 0   
     data_received..................: 18 kB   744 B/s
     data_sent......................: 9.1 kB  384 B/s
     http_req_blocked...............: avg=58.99µs  min=4µs   med=13µs  max=791µs p(90)=20.2µs  p(95)=547.99µs
     http_req_connecting............: avg=36.75µs  min=0s    med=0s    max=630µs p(90)=0s      p(95)=432.79µs
     http_req_duration..............: avg=573.59µs min=159µs med=597µs max=1ms   p(90)=702.4µs p(95)=737.8µs 
       { expected_response:true }...: avg=573.59µs min=159µs med=597µs max=1ms   p(90)=702.4µs p(95)=737.8µs 
     http_req_failed................: 0.00%   ✓ 0        ✗ 105 
     http_req_receiving.............: avg=93.16µs  min=27µs  med=98µs  max=152µs p(90)=115.2µs p(95)=123.39µs
     http_req_sending...............: avg=48.84µs  min=18µs  med=48µs  max=122µs p(90)=57.2µs  p(95)=82.79µs 
     http_req_tls_handshaking.......: avg=0s       min=0s    med=0s    max=0s    p(90)=0s      p(95)=0s      
     http_req_waiting...............: avg=431.58µs min=100µs med=445µs max=823µs p(90)=546.4µs p(95)=579.59µs
     http_reqs......................: 105     4.424716/s
     iteration_duration.............: avg=2s       min=2s    med=2s    max=2.01s p(90)=2.01s   p(95)=2.01s   
     iterations.....................: 46      1.938447/s
     vus............................: 8       min=1      max=8 
     vus_max........................: 20      min=20     max=20
```

### Golang 
Or using Goland HTTP Client with the following command:
```shell 
go run cmd/client/main.go
```

This script creates a Golang client and sends **HTTP** requests to our server.
```shell
Response status: 202 Accepted
{"Status":"202","Message":"Welcome Home"}
Response status: 202 Accepted
{"Message":"Hello Amir","Time":"2022-08-21T10:52:20.04668+04:30"}
```
