package Loader

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Port    int  `yaml:"port"`
	Mysql	MysqlConfig
}

type MysqlConfig struct{
	User string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	Server string `yaml:"server"`
	Port int `yaml:"port"`
	Database string `yaml:"database"`
}

func LoadConfig() Config {
	// 默认配置
	conf := Config{
		Port:    18848,
		Mysql: MysqlConfig{User: "root", Passwd: "", Server: "127.0.0.1", Port: 3306},
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

	fmt.Println(conf, conf.Port)
	return conf
}
