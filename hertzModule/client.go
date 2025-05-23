package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	// "github.com/cloudwego/netpoll"
)

func clientGetRequest() {
	// 客户端 默认使用 netpoll.NewDialer go标准库:standard.NewDialer()
	c, err := client.NewClient(client.WithDialer(standard.NewDialer()))
	if err != nil {
		return
	}
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	req.SetRequestURI("http://localhost:8080/hello")
	req.SetMethod("get")

	err = c.Do(context.Background(), req, resp)
	if err != nil {
		return
	}
	fmt.Printf("get response: %s\n", resp.Body())
}

// 返回结果流式处理
func respStream() {
	c, err := client.NewClient()
	if err != nil {
		return
	}
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	req.SetRequestURI("http://localhost:8080/hello")
	req.SetMethod("get")

	err = c.Do(context.Background(), req, resp)
	if err != nil {
		return
	}
	fmt.Printf("get response: %s\n", resp.Body())

	bodyStram := resp.BodyStream()
	p := make([]byte, resp.Header.ContentLength()/2)
	_, err = bodyStram.Read(p)
	if err != nil {
		fmt.Println(err.Error())
	}
	left, _ := ioutil.ReadAll(bodyStram)
	fmt.Println(string(p), string(left))

}

// 上传文件
func uploadFile() {
	c, err := client.NewClient()
	if err != nil {
		return
	}
	req, resp := &protocol.Request{}, &protocol.Response{}
	req.SetRequestURI("http://localhost:8080/singleFile")
	req.SetFile("file", "your file path")

	err = c.Do(context.Background(), req, resp)
	if err != nil {
		return
	}
	fmt.Printf("get response: %s\n", resp.Body())
}

// json格式请求
func postJson() {
	c, err := client.NewClient()
	if err != nil {
		return
	}
	req, resp := &protocol.Request{}, &protocol.Response{}

	req.Header.SetMethod(consts.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetRequestURI("http://localhost:8080/singleFile")

	data := struct {
		Json string `json:"json"`
	}{
		"test json",
	}

	jsonByte, _ := json.Marshal(data)
	req.SetBody(jsonByte)

	err = c.Do(context.Background(), req, resp)
	if err != nil {
		return
	}
	fmt.Printf("get response: %s\n", resp.Body())
}
