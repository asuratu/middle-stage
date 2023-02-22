# 环境部署


### 1. 安装依赖
``
docker-compose -f docker-compose-env.yml up -d
``

### 2. kong

#### 2.1 创建docker network
``
docker network create kong-net
``
#### 2.2 postgresql-compose.yml 文件
``
docker-compose -f postgresql-compose.yml up -d
``
#### 2.3 初始化kong 数据库
``
docker run --rm --network=kong-net -e "KONG_DATABASE=postgres" -e "KONG_PG_HOST=kong-database" -e "KONG_PG_PASSWORD=kongpass" kong/kong-gateway:3.1.1.3-alpine kong migrations bootstrap
``

#### 2.4 kong-compose.yml 文件
``
`



 



 
    
