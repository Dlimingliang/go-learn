package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type ServiceConfig struct {
	ServiceName string `mapstructure:"name"`
	ServicePort int    `mapstructure:"port"`
}

func main() {

	configFile := "20_viper/env/config-dev.yaml"
	//从环境变量读取配置,确定当前环境
	viper.AutomaticEnv()
	isProd := viper.GetBool("IS_PROD")
	if isProd {
		configFile = "20_viper/env/config-prod.yaml"
	}

	v := viper.New()
	v.SetConfigFile(configFile)

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
}
