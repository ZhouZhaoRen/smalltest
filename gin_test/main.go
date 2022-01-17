package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"smalltest/gin_test/router"
)

func leakyBucket() gin.HandlerFunc {
	return nil
}

func main() {
	// 创建一个默认的路由引擎
	r := router.InitRouter()
	r.GET("", leakyBucket(), func(context *gin.Context) {

	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8090),
		Handler: r,
	}
	s.ListenAndServe()
	//r.Run()
}
