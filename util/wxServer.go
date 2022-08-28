package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	APPID            = "wx0aa24b14c3c8143e"
	APPSECRET        = "8e87fd134fc2097fd28a5a19dfa8bc2f"
	ACCESS_TOKEN_URL = "https://api.weixin.qq.com/cgi-bin/token"
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

//
//func init() {
//
//	freshTokenTicker := time.NewTicker(4 * time.Second)
//	//requestToken()
//	if access_token == "" && ExpirationTime == 0 {
//		fmt.Println("空")
//	}
//	go func() {
//
//		for range freshTokenTicker.C {
//			fmt.Println("klkl")
//		}
//	}()
//
//}
//func main() {
//
//	////time.NewTimer时间到了,只响应一次
//	////创建一个定时器,设置时间为2s,2s后,往time通道写内容(当前时间)
//	//timer := time.NewTimer(2 * time.Second)
//	//fmt.Println("当前时间: ", time.Now())
//	//
//	////2s后,往timer.c写数据,有数据后,就可以读取
//	//t := <-timer.C //channel没有数据前后阻塞
//	//fmt.Println("t = ", t)
//	//定时2秒,2秒后产生一个事件,往channel里面写内容
//	//<-time.After(2 * time.Second)
//	//fmt.Println("时间到")
//	//timer := time.NewTicker(5 * time.Second)
//	//for {
//	//	select {
//	//	case <-timer.C:
//	//		go func() {
//	//			log.Println(time.Now())
//	//		}()
//	//	}
//	//}
//	//GetAccessToken()
//	//SendTemp()
//	for {
//
//	}
//}
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
	request, err := http.NewRequest(http.MethodGet, ACCESS_TOKEN_URL, nil)
	if err != nil {
		fmt.Println("err", err)
	}
	query := make(url.Values)
	query.Add("grant_type", "client_credential")
	query.Add("appid", APPID)
	query.Add("secret", APPSECRET)
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
		Touser:     "oKZFN5jPVELwBH5Lbnl6RW29Tx0c",
		TemplateId: "Vvgp52lxAgds5mi9sNuX8uldm96c0rgLD5idBpjxLGA",
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
	//response.
	fmt.Println(response.Body, "body")

	fmt.Println(request.URL)
}
