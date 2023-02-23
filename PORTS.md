
| 服务名称      | rpc  | api  | kq   | asynq |
|-----------|------|------|------|-------|
| job       | -    | -    | -    | 3002  |
| scheduler | -    | -    | -    | 3003  |
| user      | 5001 | 6001 | 7001 | -     |


| 依赖            | 宿主机           | 容器            |
|---------------|---------------|---------------|
| elasticsearch | 9200<br/>9300 | 9200<br/>9300 |
| kibana        | 5601          | 5601          |
| asynqmon      | 8980          | 8180          |
| mysql         | 3310          | 3306          |
| redis         | 36379         | 6379          |
| etcd          | 2379<br/>2380 | 2379<br/>2380 |
| etcd-manage   | 8280          | 8080          |