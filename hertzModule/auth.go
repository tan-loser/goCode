package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/basic_auth"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func authTest() {
	h := server.New(server.WithHostPorts(":8080"))
	h.Use(basic_auth.BasicAuth(map[string]string{
		"username1": "password1",
		"username2": "password2",
	}))

	h.GET("/get", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "HELLO HERTZ")
	})

	h.Spin()

}

// func main(){
// 	authTest()
// }
