.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo-proto && cwgo server -I ../../idl --type RPC --module github.com/cloudwego/biz-demo/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo-thrift &&cwgo server --type RPC --module gomall --service demo_thrift --idl ../../idl/echo.thrift
