package GracefulShutDown

import (
	"context"
	"fmt"
	"go-chatbot/Settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 此函数用于优雅关闭服务
func ShutDown(r *gin.Engine) {
	// 构造HTTP Server
	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d",
			Settings.Conf.Host,
			Settings.Conf.Port,
		),
		Handler: r,
	}
	// 启动Server
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			zap.L().Fatal("listen:%s\n", zap.Error(err))
		}
	}()
	// 定义退出channel
	quit := make(chan os.Signal, 1)
	// 勾取退出信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 监听信号
	<-quit
	// 日志记录
	zap.L().Info("Shutdown Server ...")
	// 设置shutdown上下文,5s超时
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	// 调用shutdown
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Shutdown", zap.Error(err))
	}
	zap.L().Info("Exiting ...")
}
