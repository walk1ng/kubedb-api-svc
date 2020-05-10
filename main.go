package main

import (
	"kubedb-api-svc/api/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong!")
	})

	my := route.Group("/my")
	{
		// 获取指定命名空间的 mysql 实例列表
		my.GET("/:ns", mysql.ListMysql)

		// 获取指定命名空间，指定名称的 mysql 实例
		my.GET("/:ns/:name", mysql.GetMysql)

		// 创建 mysql 实例在指定命名空间
		my.POST("/:ns", mysql.CreateMysql)

		// 删除指定命名空间和命名的 mysql 实例
		my.DELETE("/:ns/:name", mysql.DeleteMysql)
	}

	route.Run()
}
