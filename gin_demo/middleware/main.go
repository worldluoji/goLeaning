package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger中间件实现
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置变量example
		log.Print("before request")
		c.Set("example", "666666")

		// 请求之前

		c.Next()

		// 请求之后
		latency := time.Since(t)
		log.Println("after request", latency)

		// 访问我们发送的状态, 200, 404等等
		status := c.Writer.Status()
		log.Println("status", status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "666666"
		log.Println("request processing", example)
		c.JSON(http.StatusOK, example)
	})

	// Listen and serve on 0.0.0.0:8089
	r.Run(":8089")
}
