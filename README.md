# go_gateway_demo
===========================

###########环境依赖
go 1.16.3
Redis 
Mysql

###########部署步骤
1. 导入数据库 go_gateway

2. 启动redis 端口6379

3. main.go目录下启动:
go run main.go -configs=./conf/dev/ -endpoint=dashboard
go run main.go -configs=./conf/dev/ -endpoint=server

4. Swagger测试
http://127.0.0.1:8880/swagger/index.html#/

###########目录结构描述
├── Readme.md                   // help  
├── cert_file                   // Https证书  
├── conf/dev                    // 默认配置  
├── controller                  // 控制层  
├── dao                         // 数据访问层  
├── docs                        // Swagger目录
├── dto                         // 传输层
├── gorm                        // gorm包
├── http-proxy_middleware       // 服务端中间件
├── http_proxy_router           // 服务端路由
├── lib                         // lib包
├── log                         // 日志包
├── logs                        
├── middleware                  //中间件
├── public                      //公用函数
├── reverse_proxy               //反向代理
├── router                      //路由
├── main.go                     //main.go
└── go.mod



###########V1.0.0 框架
1. GO框架    Gin
2. Mysql     Gorm控制
3. Vue       element
4. 反向代理  reverse_proxy
