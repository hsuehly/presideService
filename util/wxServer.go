package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hsuehly/presideService/config"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var (
	access_token   string
	ExpirationTime int
)

type AccessTokenRep struct {
	Access_Token string `json:"access_token"`
	Expires_In   int    `json:"expires_in"`
}
type TempData struct {
	Value string `json:"value"`
	Color string `json:"color"`
}
type WXPushTemp struct {
	Touser     string               `json:"touser"`
	TemplateId string               `json:"template_id"`
	Url        string               `json:"url"`
	Topcolor   string               `json:"topcolor"`
	Data       map[string]*TempData `json:"data"`
}

func WxinitData() {
	//GetAccessToken()
	freshTokenTicker := time.NewTicker(7000 * time.Second)

	if access_token == "" && ExpirationTime == 0 {
		//fmt.Println("空")
		GetAccessToken()
	}
	//requestToken()

	go func() {

		for range freshTokenTicker.C {
			//fmt.Println("klkl")

			GetAccessToken()
		}
	}()
}
func GetAccessToken() {
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, "https://api.weixin.qq.com/cgi-bin/token", nil)
	if err != nil {
		fmt.Println("err", err)
	}
	query := make(url.Values)
	query.Add("grant_type", "client_credential")
	query.Add("appid", config.Configs.WxConfig.AppId)
	query.Add("secret", config.Configs.WxConfig.AppSecret)
	request.URL.RawQuery = query.Encode()
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("err", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("转换错误", err)
	}
	var req AccessTokenRep
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println("json err", err)
	}
	access_token = req.Access_Token
	ExpirationTime = req.Expires_In
	fmt.Println(access_token)
	fmt.Println(ExpirationTime)
}

func SendTemp(value map[string]*TempData) {
	//var TempD = make(map[string]*TempData)
	//TempD["name"] = &TempData{
	//	Value: "hsuehly",
	//	Color: "#173177",
	//}
	var Temp = WXPushTemp{
		Touser:     config.Configs.WxConfig.Touser,
		TemplateId: config.Configs.WxConfig.TemplateId,
		Url:        "https://www.baidu.com",
		Topcolor:   "#FF0000",
		Data:       value,
	}
	body, err := json.Marshal(&Temp)
	if err != nil {
		fmt.Println("err", err)
	}
	client := http.Client{}
	request, err := http.NewRequest(http.MethodPost, "https://api.weixin.qq.com/cgi-bin/message/template/send", bytes.NewReader(body))
	if err != nil {
		fmt.Println("err", err)
	}

	query := make(url.Values)
	query.Add("access_token", access_token)
	request.URL.RawQuery = query.Encode()
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("err", err)
	}
	defer response.Body.Close()

	fmt.Println(response.Body, "body")

	fmt.Println(request.URL)
}
