# About This

[Write a Kubernetes-ready service from zero step-by-step](https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/)を写経、改変したもの。

# Usage

## Build and Run
```
$ go run main.go -port 8000
```
or
```
$ go build
$ ./your-binary-name -port 8000
```
or
```
$ make run
```

## CURL
```
$ curl http://127.0.0.1:8000/home
```
