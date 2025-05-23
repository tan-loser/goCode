package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func engineTest(){

	fmt.Println("Engine start!")

	h := server.New(server.WithHostPorts(":8080"))

	h.GET("/get",func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{
			"message":"engine",
		})
	})

	h.Spin()

}

// func main(){
// 	engineTest()
// }