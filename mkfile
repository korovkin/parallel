travis:
	go build -o parallel cmd/*.go

build: thrift
	go build -o parallel cmd/*.go

thrift:
	thrift -r --gen go parallel.thrift
	@mv gen-go/parallel/*.go .
	@rm -rfv gen-go

