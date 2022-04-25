package Settings

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

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
}

type LogConfig struct {
	Level      string
	FileName   string
	MaxSize    int `mapstructure:"max_size"`
	MaxAge     int `mapstructure:"max_age"`
	MaxBackUps int `mapstructure:"max_backups"`
}

type MySqlConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	DBName       string
	MaxOpenConns int `mapstructure:"max_open_conns"`
	MaxIdleConns int `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
	PoolSize int
}

type ChatConfig struct {
	Host string
	Port int
	Rout string
	Key  string
}

func Init() (err error) {
	viper.SetConfigName("Config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./Configs/")
	if err = viper.ReadInConfig(); err != nil {
		return err
	}
	if err = viper.Unmarshal(&Conf); err != nil {
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.L().Info("Config Changed...", zap.String("theConfig:", fmt.Sprintf("%v", in)))
		if err = viper.Unmarshal(&Conf); err != nil {
			zap.L().Error("Config Updata Failed ...", zap.Error(err))
		}
	})
	return nil
}
