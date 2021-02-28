
get-mock:
	@go get github.com/golang/mock/mockgen@v1.5.0

mock-gen:
	@rm -rf mocks/*
	@mockgen -destination mocks/mock_writer.go -package=mocks io Writer