package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

type MysqlConfig struct {
	IP   string `mapstructure:"ip"`
	Port string `mapstructure:"port"`
}

type ServiceConfig struct {
	ServiceName string      `mapstructure:"name"`
	ServicePort int         `mapstructure:"port"`
	Mysql       MysqlConfig `mapstructure:"mysql"`
}

func main() {
	v := viper.New()
	v.SetConfigFile("20_viper/watch_config/config.yaml")

	err := v.ReadInConfig()
	if err != nil {
		return
	}
	fmt.Println(v.Get("name"))

	serviceConfig := ServiceConfig{}
	err = v.Unmarshal(&serviceConfig)
	if err != nil {
		return
	}
	fmt.Println(serviceConfig)

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config change", in.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&serviceConfig)
		fmt.Println(serviceConfig)
	})

	time.Sleep(300 * time.Second)
}
