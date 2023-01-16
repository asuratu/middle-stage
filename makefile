air := $(shell which air)
sql2pb := $(shell which sql2pb)
goctl := $(shell which goctl)
date = $(shell date "+%Y-%m-%d %H:%M:%S")

dir = ~/code/go/go-zero/MiddleStage

api:
	@cd $(dir)/app/$(m)/api && $(air) $(m).go -f etc/$(m).yaml

rpc:
	@cd $(dir)/app/$(m)/rpc && $(air) $(m).go -f etc/$(m).yaml


############################################################
# 数据库相关:
# - docker exec -it mysql mysqldump -u用户名 -p密码 数据库 > /mnt/vdb/data/mysql/test_db.sql
############################################################

db.init:
	docker exec -i mysql mysql --default-character-set=utf8mb4 -uroot -proot dev < ./proto/model/schema.sql

db.dump:
	docker exec -it mysql mysqldump --no-data --default-character-set=utf8mb4 -R -E --hex-blob -uroot -proot dev > ./tmp/backup_schema-$(date).sql

db.dump.with.data:
	docker exec -it mysql mysqldump --default-character-set=utf8mb4 -R -E --hex-blob -uroot -proot dev > ./tmp/backup.sql

db.init:
	docker exec -i mysql mysql --default-character-set=utf8mb4 -uroot -proot dev < ./proto/model/schema.sql


