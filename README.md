# MiddleStage


### 使用到的技术栈
- [x] go-zero
- [x] mysql
- [x] redis
    - [x] 队列 (asynq 做延迟队列、定时队列)
    - [x] 缓存
    - [ ] 分布式锁
- [ ] rocketmq
    - [ ] 消息队列 (前期项目简单直接使用 asynq 做消息队列)
- [x] etcd (服务注册发现)
- [x] elasticsearch
    - [x] kibana
- [ ] nacos (做配置中心, 服务发现) (前期项目简单直接使用 etcd 做服务注册发现, 文件配置)
- [ ] kong (网关), 本项目使用 nginx 做网关
- [x] sentinal (限流)
- [ ] prometheus (监控)
    - [ ] grafana (监控展示)
- [ ] dtm (分布式事务)
- [x] goframe (一些简单的工具包可以使用 goframe 提供的)


### 已完成的功能列表
- [x] 用户中心
    - [x] 手机号是否存在
    - [x] 生成图形验证码
    - [x] 发送短信验证码
    - [x] 使用手机号注册
    - [x] 使用手机号密码登录
    - [x] 根据id获取用户详情
    - [x] 获取当前用户详情
    - [x] 分页获取用户列表
    - [x] 更新用户信息
    - [x] 搜索用户
      - [x] 基于elasticsearch

