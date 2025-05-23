package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func respTest() {

	fmt.Println("resp start!")

	h := server.New(server.WithHostPorts(":8080"))

	h.GET("/hello", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{
			"message": "resp",
		})
	})
	h.GET("/get", func(c context.Context, ctx *app.RequestContext) {
		ctx.NotFound()
	})

	// json格式
	h.GET("/user", func(c context.Context, ctx *app.RequestContext) {
		ctx.Write([]byte("{'foo':'bar'}"))
		ctx.SetContentType("application/json;charset=utf-8")
	})

	// 重定向
	h.GET("/redirect", func(c context.Context, ctx *app.RequestContext) {
		ctx.Redirect(consts.StatusFound, []byte("/hello")) // 或者输入完整的域名
	})

	h.Spin()

}

// func main() {
// 	respTest()
// }
