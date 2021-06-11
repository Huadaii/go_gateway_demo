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
  
  
###########V1.0.0 程序详解  
诸多的服务通过不同的入口访问，所有的请求此时都会挂在网关服务上转发，或者是更粗暴的通过 ip:port 的形式访问服务。随着服务化的进程愈演愈烈，更多的新生服务开始部署，这时候如上的做法会让管理服务变得更加痛苦。  

隔离，即划边界  
解耦，即抽组建  
脚手架，即外层拓展  

![v2-6a64cad6674b6cc3e01adfd2225b59b3_r](https://user-images.githubusercontent.com/51690238/121619237-ddce6f00-ca9a-11eb-930a-d3c99dc160ad.jpg)  

  
  
###########V1.0.0 运行截图  
//管理员登陆  
![login](https://user-images.githubusercontent.com/51690238/121618713-deb2d100-ca99-11eb-9a41-0b827113ea33.PNG)  
//Swagger接口文档服务  
![swagger](https://user-images.githubusercontent.com/51690238/121618831-191c6e00-ca9a-11eb-8966-98a54898aedf.PNG)  
//服务列表  
![dashboard](https://user-images.githubusercontent.com/51690238/121618790-06a23480-ca9a-11eb-98e3-a3ceccf64479.PNG)  
//服务统计  
![服务统计](https://user-images.githubusercontent.com/51690238/121618876-2df90180-ca9a-11eb-889f-1174055077aa.PNG)  


