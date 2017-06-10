
build: thrift
	go build

thrift:
	thrift -r --gen go parallel.thrift
	@mv gen-go/parallel/constants.go parallel_constants.go
	@mv gen-go/parallel/parallel.go parallel_service.go
	@mv gen-go/parallel/ttypes.go parallel_types.go
	@rm -rf gen-go

