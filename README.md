# MiddleStage


### ORM
- [ ] gorm 
    - [ ] migrate
    - [ ] hook
    - [ ] cmd

### 计划使用到的技术栈
- [ ] go-zero
- [ ] mysql
    - [ ] gorm 
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