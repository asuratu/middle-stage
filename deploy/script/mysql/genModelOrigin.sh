#!/usr/bin/env bash

# 此脚本不使用自定义模板，使用goctl默认模板

# 使用方法：
# ./genModel.sh usercenter user
# ./genModel.sh usercenter user_auth
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package

# 生成的表名
tables=$2
# 表生成的genmodel目录
modeldir=../../../app/$1/rpc/model
# 是否使用缓存
isCache=$3

# 数据库配置
host=127.0.0.1
port=3310
dbname=ms_$1
username=root
passwd=PXDN93VRKUm8TeE7

echo "开始创建库：$dbname 的表：$2"

if [ "$isCache" = "true" ]; then
    goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" --style=goZero -cache=true
else
    goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" --style=goZero
fi

echo "创建库：$dbname 的表：$2 完成"


