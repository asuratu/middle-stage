#!/bin/bash
## author: "AsuraTu"
## email: "455224159@qq.com"

# 项目根目录
projectPath=$(pwd)

# 用户服务
userServer=user
userApiPath=${projectPath}/app/${userServer}/api
userRpcPath=${projectPath}/app/${userServer}/rpc

# asynq服务
asynqJobPath=${projectPath}/app/mqueue/cmd/job
asynqSchedulerPath=${projectPath}/app/mqueue/cmd/scheduler

ApiServer(){
    server=$1
    dir=$2
    air_log=${projectPath}/public/air/api/${server}.log
    cd "${dir}" || return
    air "${dir}"/"${server}".go > "${air_log}" 2>&1 &
}

RpcServer(){
    server=$1
    dir=$2
    air_log=${projectPath}/public/air/rpc/${server}.log
    cd "${dir}" || return
    air "${dir}"/"${server}".go > "${air_log}" 2>&1 &
}

AsynqServer(){
    server=$1
    dir=$2
    air_log=${projectPath}/public/air/mqueue/${server}.log
    cd "${dir}" || return
    air "${dir}"/mqueue.go > "${air_log}" 2>&1 &
}

killall air

# asynq服务
AsynqServer "job" "${asynqJobPath}"
AsynqServer "scheduler" "${asynqSchedulerPath}"

#用户服务.
RpcServer "${userServer}" "${userRpcPath}"
ApiServer "${userServer}" "${userApiPath}"
