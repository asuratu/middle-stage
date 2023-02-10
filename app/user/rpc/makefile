air := $(shell which air)
sql2pb := $(shell which sql2pb)
goctl := $(shell which goctl)
dir = ~/code/go/go-zero/MiddleStage

sql2pb:
	@$(sql2pb) -go_package ./user -host localhost -package user -password PXDN93VRKUm8TeE7 -port 3310 -schema ms_user -ignore_tables "link" -service_name usersvc -user root > usersvc.proto

zrpc:
	goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.


