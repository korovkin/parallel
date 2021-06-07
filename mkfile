
travis:
	go build cmd/parallel.go

build: thrift
	go build

thrift:
	thrift -r --gen go parallel.thrift
	@mv gen-go/parallel/*.go .
	@rm -rfv gen-go

