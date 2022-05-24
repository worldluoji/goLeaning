package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Product struct {
	Username    string    `json:"username" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	Price       int       `json:"price" binding:"gte=0"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type productHandler struct {
	sync.RWMutex
	products map[string]Product
}

func newProductHandler() *productHandler {
	return &productHandler{
		products: make(map[string]Product),
	}
}

// curl -X POST -H 'Content-Type: application/json' -d '{"username":"luoji","name":"mate50","category":"mobile","price":5000,"description":"Greate Mobile"}' http://localhost:8098/v1/products
func (u *productHandler) Create(c *gin.Context) {
	// 	log.Println(c.GetHeader("Content-Type"))  -> would print application/json
	u.Lock()
	defer u.Unlock()

	// 1. 参数解析
	var product Product
	// 通过c.ShouldBindJSON函数，将 Body 中的 JSON 格式数据解析到指定的 Struct 中
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. 参数校验
	if _, ok := u.products[product.Name]; ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("product %s already exist", product.Name)})
		return
	}
	product.CreatedAt = time.Now()

	// 3. 逻辑处理
	u.products[product.Name] = product
	log.Printf("Register product %s success", product.Name)

	// 4. 返回结果
	c.JSON(http.StatusOK, product)
}

// http://localhost:8098/v1/products/mate50
func (u *productHandler) Get(c *gin.Context) {
	u.Lock()
	defer u.Unlock()

	// c.Param获取路径参数
	product, ok := u.products[c.Param("name")]
	// log.Println(c.Param("name")) -> would print mate50
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Errorf("can not found product %s", c.Param("name"))})
		return
	}

	c.JSON(http.StatusOK, product)
}

// http://localhost:8098/v1/products/query?name=huawei
func (u *productHandler) GetQueryParam(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Errorf("failed to get param")})
		return
	}
	c.JSON(http.StatusOK, name)
}

func router() http.Handler {
	router := gin.Default()
	productHandler := newProductHandler()

	// 路由分组、中间件、认证
	v1 := router.Group("/v1")
	{
		productv1 := v1.Group("/products")
		{
			// 路由匹配
			productv1.POST("", productHandler.Create)
			productv1.GET(":name", productHandler.Get)
			productv1.GET("/query", productHandler.GetQueryParam)
		}
	}

	return router
}

func main() {
	var eg errgroup.Group

	// 一进程多端口
	insecureServer := &http.Server{
		Addr:         ":8098",
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	secureServer := &http.Server{
		Addr:         ":8543",
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	eg.Go(func() error {
		err := insecureServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	eg.Go(func() error {
		err := secureServer.ListenAndServeTLS("server.pem", "server.key")
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
