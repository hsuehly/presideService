package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App App `yaml:"app" json:"app"`
	Log Log `yaml:"log" json:"log"`
}
type App struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

type Log struct {
	Suffix  string `yaml:"suffix" json:"suffix"`
	MaxSize int    `yaml:"maxSize" json:"maxsize"`
}

func InitConfigData() {
	//导入配置文件
	yamlFile, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	var _config *Config
	//将配置文件读取到结构体中
	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}
	username := _config.App.Username
	fmt.Println("username:" + username)
	fmt.Printf("config.log: %#v\n", _config.Log)

}
