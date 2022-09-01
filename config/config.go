package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Dev Dev `yaml:"dev" json:"dev"`
	Pro Pro `yaml:"pro" json:"pro"`
}
type Dev struct {
	Mysql Mysql `yaml:"mysql" json:"mysql"`
}
type Mysql struct {
	Host     int    `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

type Pro struct {
}

func InitConfigData() {
	//导入配置文件
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	var _config *Config
	//将配置文件读取到结构体中
	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}
	username := _config.Dev.Mysql.Port
	fmt.Println("username:", username)

}
