air := $(shell which air)
sql2pb := $(shell which sql2pb)
goctl := $(shell which goctl)
dir = ~/code/go/go-zero/MiddleStage

api:
	@cd $(dir)/app/$(m)/api && $(air) $(m).go -f etc/$(m).yaml

rpc:
	@cd $(dir)/app/$(m)/rpc && $(air) $(m).go -f etc/$(m).yaml

sql2pb:
	@$(sql2pb) -go_package ./pb -host localhost -package user -password '' -port 3306 -schema stage_user -service_name user -user root > user.proto





