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
	v := viper.New()
	v.SetConfigFile("20_viper/config.yaml")

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
