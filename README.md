# MiddleStage


### ORM
- [ ] gorm 
    - [ ] migrate
    - [ ] hook
    - [ ] cmd

### 计划使用到的技术栈
- [ ] go-zero
- [ ] mysql
- [ ] redis
    - [ ] 队列 (asynq 做延迟队列、定时队列)
    - [ ] 缓存
    - [ ] 分布式锁
- [ ] rocketmq
    - [ ] 消息队列 (前期项目简单直接使用 asynq 做消息队列)
- [ ] etcd (服务注册发现)
- [ ] elasticsearch
    - [ ] kibana
- [ ] nacos (做配置中心, 服务发现) (前期项目简单直接使用 etcd 做服务注册发现, 文件配置)
- [ ] kong (网关)
- [ ] sentinal (限流)
- [ ] prometheus (监控)
    - [ ] grafana (监控展示)
- [ ] dtm (分布式事务)
- [ ] goframe (一些简单的工具包可以使用 goframe 提供的)


### 未完成的功能列表
- [ ] 用户中心
    - [ ] 注册
    - [ ] 登录
    - [ ] 获取用户信息
- [ ] 搜索
    - [ ] 基于elasticsearch


### 需要注意的是 elasticsearch 第一次用的时候，需要初始化密码 执行下面的操作
```bash
➜ ~ docker exec -it es /bin/bash
[root@04b37b58f104 elasticsearch]# sh /usr/share/elasticsearch/bin/elasticsearch-setup-passwords auto
Initiating the setup of passwords for reserved users elastic,apm_system,kibana,kibana_system,logstash_system,beats_system,remote_monitoring_user.
The passwords will be randomly generated and printed to the console.
Please confirm that you would like to continue [y/N]y


Changed password for user apm_system
PASSWORD apm_system = 00UcLVZKBIbmVq2LDeNn

Changed password for user kibana_system
PASSWORD kibana_system = LUyQFglN7NorJ5C3aORq

Changed password for user kibana
PASSWORD kibana = LUyQFglN7NorJ5C3aORq

Changed password for user logstash_system
PASSWORD logstash_system = NvtUm3yA6sYjcXSF7Dyj

Changed password for user beats_system
PASSWORD beats_system = cCgJnQc7yP68dfSXsHLH

Changed password for user remote_monitoring_user
PASSWORD remote_monitoring_user = 2you6KdlsIsfS4M2CKd5

Changed password for user elastic
PASSWORD elastic = YGfgsppPByhB3IjFtFu4
```
### 如果报错
```bash
It doesn't look like the X-Pack security feature is enabled on this Elasticsearch node.
```
### 需要修改配置文件 elasticseach.yml 后, 重启容器
```yaml
http.host: 0.0.0.0
xpack.security.enabled: true
xpack.security.transport.ssl.enabled: true
```
### 看到上面最后行
```
Changed password for user elastic
PASSWORD elastic = YGfgsppPByhB3IjFtFu4
```
### 得到elastic 的帐号: elastic ,密码: YGfgsppPByhB3IjFtFu4 将这个填入search/rpc/search.yaml 文件中