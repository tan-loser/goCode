/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/basic_auth"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func myRecoveryHandler(c context.Context, ctx *app.RequestContext, err interface{}, stack []byte) {
	hlog.SystemLogger().CtxErrorf(c, "[Recovery] error=%v\nstack=%s", err, stack)
	hlog.SystemLogger().Infof("Client: %s", ctx.Request.Header.UserAgent())
	ctx.AbortWithStatus(consts.StatusInternalServerError)
}

func main1() {
	// server.Default() creates a Hertz with recovery middleware.
	// If you need a pure hertz, you can use server.New()
	//  h := server.Default()
	// Server 默认使用netpoll,不支持window(23年的视频) netpoll.NewTransporter
	h := server.New(server.WithHostPorts(":8080"), server.WithAltTransport(standard.NewTransporter))

	// 处理异常的中间件 当返回panic时，返回500并打印错误信息
	// h.Use(recovery.Recovery())
	// 使用自定义的处理器
	h.Use(recovery.Recovery(recovery.WithRecoveryHandler(myRecoveryHandler)))
	h.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
		// 使用panic测试recovery
		// panic("test")

		//  c.String(consts.StatusOK, "Hello hertz!")
		c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	//路由分组
	v1 := h.Group("/v1")
	v1.GET("/test", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "v1 test"})
	})

	// http://localhost:8080/v2/test
	// 分组可以 中间件 鉴权
	v2 := h.Group("/v2", basic_auth.BasicAuth(map[string]string{"test": "test"}))
	// 或者使用Use
	v2.Use(basic_auth.BasicAuth(map[string]string{"test": "test"}))
	v2.GET("/test", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "v2 test"})
	})

	h.Spin()
}
