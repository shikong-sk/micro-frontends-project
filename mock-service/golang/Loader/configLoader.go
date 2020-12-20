package Loader

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// 全局配置
type Config struct {
	Port    int           `yaml:"port"`
	Mysql   MysqlConfig   `yaml:"mysql"`
	Session SessionConfig `yaml:"session"`
}

// Mysql配置
type MysqlConfig struct {
	User        string `yaml:"user"`
	Passwd      string `yaml:"passwd"`
	Server      string `yaml:"server"`
	Port        int    `yaml:"port"`
	Database    string `yaml:"database"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
	MaxIdleConn int    `yaml:"maxIdleConn"`
}

// Session 配置
type SessionConfig struct {
	Secret string `yaml:"secret"`
}

func LoadConfig() Config {
	// 默认配置
	conf := Config{
		Port: 18848,
		Mysql: MysqlConfig{
			User:        "root",
			Passwd:      "",
			Server:      "127.0.0.1",
			Port:        3306,
			MaxOpenConn: 10,
			MaxIdleConn: 5,
		},
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

	if conf.Port <= 0 || conf.Port > 65535 {
		panic("端口不合法")
	}

	fmt.Println(conf, conf.Port)
	return conf
}
