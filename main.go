package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
)

var (
	token string = ""
	url   string = "https://oapi.dingtalk.com/robot/send?access_token=" + token
)

type RebotRequest struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

type RebotResponse struct {
	ErrMsg  string `json:"errmsg"`  //响应消息
	Errcode int    `json:"errcode"` //响应状态码
}

func main() {
	req := new(RebotRequest)
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
	/* 	str := (*string)(unsafe.Pointer(&respBytes))
	   	fmt.Println(*str) */

	var reResp RebotResponse
	err = json.Unmarshal(respBytes, &reResp)
	if err != nil {
		beego.Error("解析错误", err)
		return
	}
	if reResp.Errcode == 0 && reResp.ErrMsg == "ok" {
		beego.Info("消息发送成功")
	} else {
		beego.Error("消息发送错误", err)
	}
}
