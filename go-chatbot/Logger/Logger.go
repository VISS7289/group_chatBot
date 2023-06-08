package Logger

import (
	"go-chatbot/Settings"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 初始化一个Zap日志记录器
func Init(cfg Settings.LogConfig, mode string) (err error) {
	// 根据日志配置参数cfg
	WriteSyncer := getLoggerWriter(cfg.FileName,
		cfg.MaxSize,
		cfg.MaxBackUps,
		cfg.MaxAge,
	)
	// 使用getEncoder方法获取一个Zap日志记录器的编码器
	Encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(Settings.Conf.LogConfig.Level))
	if err != nil {
		return err
	}
	var core zapcore.Core
	// 如果日志记录器的模式为"dev"，则使用NewConsoleEncoder方法创建一个控制台编码器，
	// 并使用NewCore方法创建一个Zap核心记录器实例，该实例将日志输出到标准输出流中
	if mode == "dev" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)
	} else {
		// 如果日志记录器的模式不是"dev"，则使用NewCore方法创建一个Zap核心记录器实例，该实例将日志输出到文件中
		core = zapcore.NewCore(Encoder, WriteSyncer, zapcore.DebugLevel)
	}
	// 创建一个Zap日志记录器实例
	lg := zap.New(core, zap.AddCaller())
	// 将创建的Zap日志记录器实例替换为全局日志记录器
	zap.ReplaceGlobals(lg)
	return nil
}

// 创建一个io.Writer接口，用于向指定的文件中写入日志
func getLoggerWriter(filename string, maxSize int, maxBackups int, maxAge int) zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   false,
	}
	return zapcore.AddSync(lumberjackLogger)
}

// 获取Zap日志记录器的编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// GinLogger 接收gin框架默认的日志
// 记录HTTP请求的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取当前时间，作为请求处理的起始时间
		start := time.Now()
		// 并保存到变量path和query中
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		// 等待请求处理完成
		c.Next()
		// 取请求处理的耗时，并保存到变量cost中
		cost := time.Since(start)
		// 记录请求处理的日志信息，
		//包括HTTP请求的URL路径、HTTP请求的方法、HTTP请求的查询字符串、
		//HTTP请求的IP地址、HTTP请求的User-Agent、HTTP请求处理的错误信息、
		//HTTP请求处理的耗时和HTTP响应的状态码等
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
// 捕获和处理在HTTP请求处理过程中发生的panic异常
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 判断panic异常的类型和内容，如果是*net.OpError类型的错误，
				//并且其错误信息中包含"broken pipe"或"connection reset by peer"字符串，
				//则将异常标记为"broken pipe"。
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				// 如果异常被标记为"broken pipe"，则使用Zap日志记录器的Error方法记录请求处理的错误信息，
				//包括HTTP请求的URL路径、HTTP请求处理的错误信息和HTTP请求的原始信息，
				//并使用c.Error方法将错误信息保存到Gin框架的上下文中，并使用c.Abort方法停止该请求的处理，并返回结果。
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				// 如果异常没有被标记为"broken pipe"，则根据是否需要记录堆栈信息，
				//使用Zap日志记录器的Error方法记录请求处理的错误信息，
				//包括HTTP请求的URL路径、HTTP请求处理的错误信息、
				//HTTP请求的原始信息和异常的堆栈信息（如果需要记录堆栈信息）。
				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				// 返回HTTP 500错误
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		// 继续处理后续的请求处理流程
		c.Next()
	}
}
