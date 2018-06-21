package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"unsafe"
)

var (
	token string = ""
	url   string = "https://oapi.dingtalk.com/robot/send?access_token=" + token
)

type RequestRebot struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

func main() {
	req := new(RequestRebot)
	req.MsgType = "text"
	req.Text.Content = "哈哈哈哈哈哈哈"
	req.At.AtMobiles = []string{""}
	req.At.IsAtAll = false

	bytesData, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)

	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		beego.Info("Error", err)
		return
	}
	// 必须将字符集编码设置成UTF-8
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
}
