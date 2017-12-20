APP?=go-learning-kubernetes
PORT?=8000

clean:
	rm -f ${APP}

build: clean
	go build -o ${APP}

run: build
	./${APP} -port=${PORT} 

test:
	go test -v -race ./...