package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()

	// 认证, 访问时会弹出认证框
	router.Use(gin.BasicAuth(gin.Accounts{"foo": "bar", "colin": "colin404"}))

	// RequestID https://github.com/gin-contrib/requestid
	// 响应头中会返回X-Request-Id: test
	router.Use(
		requestid.New(
			requestid.WithGenerator(func() string {
				return "test"
			}),
			requestid.WithCustomHeaderStrKey("X-Request-ID"),
		),
	)

	// 跨域 https://github.com/gin-contrib/cors
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - "GET", "POST", "PUT", "DELETE", "OPTIONS" methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/test", func(c *gin.Context) {
		log.Println("request processing")
		c.JSON(http.StatusOK, "success")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	router.Run(":8089")
}
