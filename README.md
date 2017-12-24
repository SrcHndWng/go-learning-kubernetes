# About This

[Write a Kubernetes-ready service from zero step-by-step](https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/)を写経、改変したもの。

# Usage

## Build and Run

### go command
```
$ go run main.go -port 8000
```
or
```
$ go build
$ ./your-binary-name -port 8000
```

### with docker, make

First, start docker.

```
$ make run
```

## Test
```
$ go test -v ./...
```

## CURL

### local

```
$ curl -i http://127.0.0.1:8000/home
$ curl -i http://127.0.0.1:8000/healthz
$ curl -i http://127.0.0.1:8000/readyz
```

### Docker

```
$ curl -i http://192.168.99.100:8000/home
```
