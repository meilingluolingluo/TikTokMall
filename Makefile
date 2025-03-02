.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo-proto && cwgo server -I ../../idl --type RPC --module github.com/cloudwego/biz-demo/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo-thrift &&cwgo server --type RPC --module gomall --service demo_thrift --idl ../../idl/echo.thrift

.PHONY: demo-link-fix
demo-link-fix:
	@cd demo/demo-proto && ln -sf ../../idl/echo.proto .
	@cd demo/demo-thrift && ln -sf ../../idl/echo.thrift .

.PHONY: demo-clean
demo-clean:
	@cd demo/demo-proto && rm -rf echo.proto
	@cd demo/demo-thrift && rm -rf echo.thrift

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend --module github.com/cloudwego/biz-demo/gomall/app/frontend -I ../../idl