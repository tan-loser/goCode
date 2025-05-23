package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// 静态路由 > 命名参数路由 > 统配参数路由
func staticRou() {
	fmt.Println("路由类型")

	h := server.New()

	// 静态路由
	h.GET("/get", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{
			"message": "router",
		})
	})

	// 命名参数路由 :name 通过RequestContext.Param("name")获取
	h.GET("/get/:name", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{
			"message": "router " + ctx.Param("name"),
		})
	})

	// 通配符参数路由 *path
	h.GET("/get/*path", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{
			"message": "router " + ctx.Param("path"),
		})
	})

	h.Spin()
}

// func main() {
// 	staticRou()
// }
