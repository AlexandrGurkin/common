VERSION := $(shell git describe --tags 2> /dev/null || echo no-tag)
BRANCH := $(shell git symbolic-ref -q --short HEAD)
COMMIT := $(shell git rev-parse HEAD)

get-mock:
	@go get github.com/golang/mock/mockgen@v1.5.0

mock-gen:
	@rm -rf mocks/mock_writer.go
	@mockgen -destination mocks/mock_writer.go -package=mocks github.com/AlexandrGurkin/common/xlog WriteSyncer

tools:
	@go get golang.org/x/tools/cmd/benchcmp

logs-bench: tools
	@go test ./xlog/xzerolog/ -bench=. -benchmem > ./xlog/bench/zerolog_$(VERSION).txt
	@go test ./xlog/xlogrus/ -bench=. -benchmem > ./xlog/bench/ruslog_$(VERSION).txt
	@go test ./xlog/xzap/ -bench=. -benchmem > ./xlog/bench/zaplog_$(VERSION).txt
	@benchcmp ./xlog/bench/zaplog_$(VERSION).txt ./xlog/bench/zerolog_$(VERSION).txt

test-all:
	@go test ./...
