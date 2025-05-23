package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
)

// 解决跨域问题 使用中间件 cors
// 安装：go get github.com/hertz-contrib/cors
func corsTest() {
	h := server.Default()
	h.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://foo.com"},
		AllowMethods: []string{"put", "patch"},
		// 通过方法判断是否通过
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
	}))

	h.GET("/get", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{
			"message": "engine",
		})
	})

	h.Spin()
}

// JWT
// go get github.com/hertz-contrib/jwt
// type login struct {
// 	username string `form:"username" json:"username"`
// 	password string `form:"password" json:"password"`
// }

// func JWTtest() {
// 	h := server.New()
// 	h.use()

// }

// func main() {
// 	corsTest()
// }
