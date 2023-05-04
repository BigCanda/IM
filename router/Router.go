// Package router 路由层
package router

import (
	"Im/docs"
	"Im/service"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	// 创建默认的gin.Engine对象
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	// 将swagger映射到对应路径
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 将相应方法映射到对应路径
	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserList", service.GetUserList)
	r.GET("/product", service.GetProductList)
	r.GET("/user/register", service.Register)
	return r
}
