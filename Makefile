.PHONY: genidl
genidl:
	go install github.com/cloudwego/kitex/tool/cmd/kitex
	go install github.com/cloudwego/thriftgo
	cd api/kitex && rm -rf kitex_gen && kitex -combine-service -module Oasyce-backend oasyce.thrift


