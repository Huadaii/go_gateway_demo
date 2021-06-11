package http_proxy_middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gateway_demo/dao"
	"go_gateway_demo/middleware"
	"go_gateway_demo/public"
)

//匹配接入方式，基于请求信息
//请求信息与服务列表做匹配关系，匹配到所需要的服务
//用于负载均衡 反向代理 权限校验

func HTTPAccessModeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		service, err := dao.ServiceManagerHandler.HTTPAccessMode(c)
		if err != nil {
			middleware.ResponseError(c, 1001, err)
			c.Abort()
			return
		}
		fmt.Println("matched service", public.Obj2Json(service))
		//设置上下文信息
		c.Set("service", service)
		c.Next()
	}
}
