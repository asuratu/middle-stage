### 用户服务

#### api 为对外提供的 restful api 
#### rpc 为供内部服务调用的 rpc 接口

### 创建 api 文件
```shell 生成接口
goctl api -o user.api
```

### 创建 proto 文件
```shell 生成接口
goctl rpc -o user.proto
```

### 生成user api服务
```
goctl api go -api user.api -dir .

goctl rpc proto -src user.proto -dir .
```