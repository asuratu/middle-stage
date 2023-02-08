air := $(shell which air)
sql2pb := $(shell which sql2pb)
goctl := $(shell which goctl)
dir = ~/code/go/go-zero/MiddleStage

gen-model:
	@$(goctl) model mysql datasource -url="root:PXDN93VRKUm8TeE7@tcp(localhost:3310)/ms_user" -table="$(m)"  -dir="model" -cache=true --style=goZero --home ../../../tpl

sql2pb:
	@$(sql2pb) -go_package ./user -host localhost -package user -password PXDN93VRKUm8TeE7 -port 3310 -schema ms_user -table user -service_name user -user root > user.proto

zrpc:
	goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.


