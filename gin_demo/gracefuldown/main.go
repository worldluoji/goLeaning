package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

/*
* graceful shutdown 的一般思路
* 注意，sending os.Interrupt 在windows下没有实现，该程序需要在linux、macos下运行
 */
func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8082",
		Handler: router,
	}

	go func() {
		// 将服务在 goroutine 中启动
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	/*
	* The only signal values guaranteed to be present in the os package on all systems are os.
	* Interrupt (send the process an interrupt) and os.Kill (force the process to exit).
	* On Windows, sending os.Interrupt to a process with os.Process.Signal is not implemented;
	* it will return an error instead of sending a signa
	 */
	signal.Notify(quit, os.Interrupt)
	<-quit // 阻塞等待接收 channel 数据
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 5s 缓冲时间处理已有请求
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil { // 调用 net/http 包提供的优雅关闭函数：Shutdown
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
