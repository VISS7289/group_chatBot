package Settings

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

// 项目总设置
type AppConfig struct {
	Name        string
	Mode        string
	Host        string
	StartTime   string `mapstructure:"start_time"`
	Port        int
	MachingID   int64 `mapstructure:"maching_id"`
	LogConfig   `mapstructure:"log"`
	MySqlConfig `mapstructure:"mysql"`
	RedisConfig `mapstructure:"redis"`
	ChatConfig  `mapstructure:"chat"`
	EmailConfig `mapstructure:"email"`
}

// 日志设置
type LogConfig struct {
	Level      string
	FileName   string
	MaxSize    int `mapstructure:"max_size"`
	MaxAge     int `mapstructure:"max_age"`
	MaxBackUps int `mapstructure:"max_backups"`
}

// mysql设置
type MySqlConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	DBName       string
	MaxOpenConns int `mapstructure:"max_open_conns"`
	MaxIdleConns int `mapstructure:"max_idle_conns"`
}

// redis设置
type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
	PoolSize int
}

// 聊天设置
type ChatConfig struct {
	Host string
	Port int
	Rout string
	Key  string
}

// 邮箱设置
type EmailConfig struct {
	From     string
	Password string
	Host     string
	Port     int
}

// Init 函数用于初始化应用程序的配置
func Init() (err error) {
	// 设置配置文件名、文件类型和文件路径
	viper.SetConfigName("Config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./Configs/")
	// 读取配置文件
	if err = viper.ReadInConfig(); err != nil {
		return err
	}
	// 将配置文件中的配置信息反序列化到 Conf 对象中
	if err = viper.Unmarshal(&Conf); err != nil {
		return err
	}
	// 监听配置文件的变化，并在配置文件发生变化时自动更新配置信息
	viper.WatchConfig()
	// 当配置文件发生变化时执行以下操作
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.L().Info("Config Changed...", zap.String("theConfig:", fmt.Sprintf("%v", in)))
		// 将配置文件中的配置信息反序列化到 Conf 对象中
		if err = viper.Unmarshal(&Conf); err != nil {
			zap.L().Error("Config Updata Failed ...", zap.Error(err))
		}
	})
	return nil
}
