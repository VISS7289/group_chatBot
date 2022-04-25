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

func ShutDown(r *gin.Engine) {
	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d",
			Settings.Conf.Host,
			Settings.Conf.Port,
		),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			zap.L().Fatal("listen:%s\n", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Shutdown", zap.Error(err))
	}
	zap.L().Info("Exiting ...")
}
