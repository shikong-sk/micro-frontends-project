package Loader

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Version string `yaml:"version"`
	Port    int  `yaml:"port"`
}

func LoadConfig() Config {
	// 默认配置
	conf := Config{
		Version: "dev",
		Port:    18848,
	}

	// 尝试读取配置文件
	var c, err = ioutil.ReadFile("./config.yml")
	if err == nil {
		fmt.Println(c, "\n", string(c))
		err = yaml.Unmarshal(c, &conf)
		if err != nil {
			panic("配置文件解析失败, 请检查格式是否正确")
		}
	} else {
		fmt.Println("未读取到配置文件, 使用默认配置启动")
	}

	if conf.Port <= 0 || conf.Port > 65535{
		panic("端口不合法")
	}

	fmt.Println(conf, conf.Version, conf.Port)
	return conf
}
