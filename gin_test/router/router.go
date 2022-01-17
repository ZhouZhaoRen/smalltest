package router

import (
	"github.com/gin-gonic/gin"
	"smalltest/gin_test/service"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// --------------test------------
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	//---------------data-------------
	r.GET("/getData", service.GetData)

	return r
}
