host=localhost
port=8000

run:
	go run ./cmd/server/main.go -host=${host} -port=${port}

build: 
	go build ./cmd/server/main.go