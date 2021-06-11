package http_proxy_router

import (
	"github.com/gin-gonic/gin"
	http_proxy_middleware "go_gateway_demo/http-proxy_middleware"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	//todo 优化点1
	//router := gin.Default()
	router := gin.New()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Use(http_proxy_middleware.HTTPAccessModeMiddleware())
	router.Use(http_proxy_middleware.HTTPFlowCountMiddleware())
	router.Use(http_proxy_middleware.HTTPFlowLimitMiddleware())
	router.Use(http_proxy_middleware.HTTPReverseProxyMiddleware())
	router.Use(http_proxy_middleware.HTTPHeaderTransferMiddleware())
	router.Use(http_proxy_middleware.HTTPStripUriMiddleware())
	router.Use(http_proxy_middleware.HTTPUrlRewriteMiddleware())
	router.Use(http_proxy_middleware.HTTPWhiteListMiddleware())
	router.Use(http_proxy_middleware.HTTPBlackListMiddleware())
	return router
}
