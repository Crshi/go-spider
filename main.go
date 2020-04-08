package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	// 创建一个http.Server实例，指定端口与Handler。 声明一个processed chan，其用来保证服务优雅的终止后再退出主goroutine。
	// 新启一个goroutine，其会监听os.Interrupt信号，一旦服务被中断即调用服务的Shutdown方法，
	// 确保活跃连接的正常返回（本代码使用的Context超时时间为3s，大于服务Handler的处理时间，所以不会超时）。
	// 处理完成后，关闭processed通道，最后主goroutine退出。

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
