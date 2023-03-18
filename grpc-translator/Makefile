all: generate lint dep build

generate: buf_generate

buf_generate:
	buf mod update
	buf generate

lint:
	#golangci-lint run --enable-all --verbose
	golangci-lint run --disable-all --verbose
dep:
	go mod download

build:
	go build

clean:
	rm grpc-translator