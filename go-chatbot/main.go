package main

import (
	"fmt"
	"go-chatbot/Dao/Mysql"
	"go-chatbot/Logger"
	"go-chatbot/Pkg/GracefulShutDown"
	"go-chatbot/Pkg/SnowFlake"
	"go-chatbot/Routs"
	"go-chatbot/Settings"

	"go.uber.org/zap"
)

func main() {
	//加载配置文件
	if err := Settings.Init(); err != nil {
		fmt.Printf("InitSetting Error:%v \n", err)
		return
	}
	//初始化日志
	if err := Logger.Init(Settings.Conf.LogConfig, Settings.Conf.Mode); err != nil {
		fmt.Printf("LoggerInit Error:%v \n", err)
		return
	}
	defer func(l *zap.Logger) {
		_ = l.Sync()
	}(zap.L())
	zap.L().Debug("LoggerInit Success...")
	//初始化mysql连接
	if err := Mysql.Init(Settings.Conf.MySqlConfig); err != nil {
		zap.L().Error("MysqlInit Error:", zap.Error(err))
		return
	}
	defer Mysql.Close()
	zap.L().Debug("MysqlInit Success...")
	//初始化redis连接
	//if err := Redis.Init(); err != nil {
	//zap.L().Error("RedisInit Error", zap.Error(err))
	//return
	//}
	//defer Redis.Close()
	//zap.L().Debug("RedisInit Success...")
	//启动雪花算法生成用户ID
	if err := SnowFlake.Init(Settings.Conf.StartTime, Settings.Conf.MachingID); err != nil {
		zap.L().Error("SnowFlakeInit Error:", zap.Error(err))
	}
	zap.L().Debug("SnowFlakeInit Success...")
	//注册路由
	r := Routs.Init(Settings.Conf.Mode)
	//启动服务(优雅关机)
	GracefulShutDown.ShutDown(r)
}
