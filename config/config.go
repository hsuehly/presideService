package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Mysql    Mysql    `yaml:"mysql" json:"mysql"`
	WxConfig WxConfig `yaml:"wxConfig" json:"wxconfig"`
}

var Configs Config

type Mysql struct {
	Database string `yaml:"database" json:"database"`
	Host     string `yaml:"host" json:"host"`
	Port     string `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}
type WxConfig struct {
	AppId      string `yaml:"APPID" json:"appId"`
	AppSecret  string `yaml:"APPSECRET" json:"appSecret"`
	Touser     string `yaml:"Touser" json:"touser"`
	TemplateId string `yaml:"TemplateId" json:"templateId"`
}

func init() {
	//导入配置文件
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println(err.Error(), "err")
	}
	//var _config *Config
	//将配置文件读取到结构体中
	err = yaml.Unmarshal(yamlFile, &Configs)
	if err != nil {
		fmt.Println(err.Error(), "err2")
	}

}
