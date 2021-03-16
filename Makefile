
get-mock:
	@go get github.com/golang/mock/mockgen@v1.5.0

mock-gen:
	@rm -rf mocks/*
	@mockgen -destination mocks/mock_writer.go -package=mocks io Writer

tools:
	@go get golang.org/x/tools/cmd/benchcmp

logs-bench: tools
	@go test ./xlog/xzerolog/ -bench=. -benchmem > zlog.txt
	@go test ./xlog/xlogrus/ -bench=. -benchmem > rlog.txt
	@benchcmp zlog.txt rlog.txt
