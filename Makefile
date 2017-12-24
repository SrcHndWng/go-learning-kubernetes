APP?=go-learning-kubernetes
PORT?=8000
RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

clean:
	rm -f ${APP}

build: clean
	go build \
		-ldflags "-s -w \
		-X main.release=${RELEASE} \
		-X main.commit=${COMMIT} \
		-X main.buildTime=${BUILD_TIME}" \
		-o ${APP}

run: build
	./${APP} -port=${PORT} 

test:
	go test -v -race ./...