1. 编写IDL文件

2. 生成代码
```sh
kitex -module example_shop idl/item.thrift

kitex -module example_shop idl/stock.thrift
```
3. 创建rpc/item rpc/stock文件夹
再分别进入各自的目录中，执行如下命令生成代码：
```sh
// item 目录下执行
kitex -module example_shop -service example.shop.item -use example_shop/kitex_gen ../../idl/item.thrift

// stock 目录下执行
kitex -module example_shop -service example.shop.stock -use example_shop/kitex_gen ../../idl/stock.thrift
```
kitex 默认会将代码生成到执行命令的目录下，kitex 的命令中：
-module 参数表明生成代码的 go mod 中的 module name，在本例中为 example_shop
-service 参数表明我们要生成脚手架代码，后面紧跟的 example.shop.item 或 example.shop.stock 为该服务的名字。
-use 参数表示让 kitex 不生成 kitex_gen 目录，而使用该选项给出的 import path。在本例中因为第一次已经生成 kitex_gen 目录了，后面都可以复用。
最后一个参数则为该服务的 IDL 文件

4. 拉取依赖
```sh
go mod tidy
```
若想要升级 kitex 版本，执行 go get -v github.com/cloudwego/kitex@latest

5. 编写商品服务逻辑
rpc/item/handler.go 文件的 GetItem 函数就对应了我们之前在 item.thrift IDL 中定义的 GetItem 方法。
添加逻辑:
```go
    resp = item.NewGetItemResp()
	resp.Item = item.NewItem()
	resp.Item.Id = req.GetId()
	resp.Item.Title = "Kitex"
	resp.Item.Description = "Kitex is an excellent framework!"
```
main.go 中的代码很简单，即使用 kitex 生成的代码创建一个 server 服务端，并调用其 Run 方法开始运行。
通常使用 main.go 进行一些项目初始化，如加载配置等。

6. 运行商品逻辑
在 rpc/item/build.sh 主要做了以下事情：
定义了一个变量 RUN_NAME，用于指定生成的可执行文件的名称，值为我们在 IDL 中指定的 namespace。本例中为 example.shop.item
创建 output 目录，此后的编译出的二进制文件放在 output/bin 下。同时将 script 目录下的项目启动脚本复制进去
根据环境变量 IS_SYSTEM_TEST_ENV 的值判断生成普通可执行文件或测试可执行文件。值为 1 则代表使用 go test -c 生成测试文件，否则正常使用 go build 命令编译。
直接执行 **`sh build.sh`** 即可编译项目。
编译成功后，生成 output 目录：
output
├── bin // 存放二进制可执行文件
│   └── example.shop.item
└── bootstrap.sh // 运行文件的脚本
执行 **`sh output/bootstrap.sh`** 即可启动编译后的二进制文件。

6.1 window需要设置wsl,linux环境才能运行sh命令
window系统 搜索 -> powershell(管理员) -> `wsl --install` -> 设置 Linux 用户名和密码(gt:gt) -> 命令行输入wsl进入linux系统
[window安装wsl](https://learn.microsoft.com/zh-cn/windows/wsl/setup/environment#set-up-your-linux-username-and-password)
注意：window系统的$PATH在linux中也有，但是执行go命令不能运行，需要加.exe后缀才能运行，使用需要在build.sh中的go命令后面加.exe后缀才能运行，使用需要在build

7. 运行API服务
有了商品服务后，接下来就让我们编写 API 服务用于调用刚刚运行起来的商品服务，并对外暴露 HTTP 接口。直接调用商品服务的接口会报错
在api文件夹下编写main.go
- 创建 itemservice.Client 
```go
import (
	"example_shop/kitex_gen/example/shop/item/itemservice"
	"github.com/cloudwego/kitex/client"
	...
)
...
c, err := itemservice.NewClient("example.shop.item", client.WithHostPorts("0.0.0.0:8888"))
if err != nil {
	log.Fatal(err)
}

```
- 并调用 GetItem() 方法返回，并返回前端展示
```go
import "example_shop/kitex_gen/example/shop/item"
...
req := item.NewGetItemReq()
req.Id = 1024
resp, err := cli.GetItem(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
if err != nil {
  log.Fatal(err)
}
```
其第一个参数为 context.Context，通过通常用其传递信息或者控制本次调用的一些行为，你可以在后续章节中找到如何使用它。
其第二个参数为本次调用的请求参数。
其第三个参数为本次调用的 options ，Kitex 提供了一种 callopt 机制，顾名思义——调用参数 ，有别于创建 client 时传入的参数，这里传入的参数仅对此次生效。 此处的 callopt.WithRPCTimeout 用于指定此次调用的超时（通常不需要指定，此处仅作演示之用）

另启一个终端，执行 go run . 命令即可启动 API 服务，监听 8889 端口，
请求 localhost:8889/api/item 即可发起 RPC 调用商品服务提供的 GetItem 接口，并获取到响应结果。

8. 在商品服务中调用库存服务

8.1 修改库存服务，用于商品服务的调用
- rpc/stock/handler.go 的 GetItem() 方法中添加如下代码,返回库存方法返回结果
```go
    resp = stock.NewGetItemStockResp()
    resp.Stock = req.GetItemId()
```
- rpc/stock/main.go 的 main() 中设置端口（默认8888）
```go
    addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")
    svr := stock.NewServer(new(StockServiceImpl), server.WithServiceAddr(addr))
```
使用 go run . 或者按照商品服务启动

8.2 补充商品服务（rpc/item/handler.go）
- 在商品服务(ItemServiceImpl)结构体中添加库存服务客户端变量(stockservice.Client)
```go
type ItemServiceImpl struct{
  	stockCli stockservice.Client
}
```
- 在GetItem() 方法中添加 初始化库存请求，并请求获取返回结果 并包装在商品的返回结果中 resp.Item.Stock
```go
    stockReq := stock.NewGetItemStockReq()
    stockReq.ItemId = req.GetId()
    stockResp, err := s.stockCli.GetItemStock(context.Background(), stockReq)
    if err != nil {
       log.Println(err)
       stockResp.Stock = 0
    }
    resp.Item.Stock = stockResp.GetStock()
```

- 新增库存客户端的方法，用于在 rpc/item/main.go中设置库存客户端，并赋值给商品服务结构体
```go
// 传入库存服务地址，返回库存服务客户端
func NewStockClient(addr string) (stockservice.Client, error) {
    return stockservice.NewClient("example.shop.stock", client.WithHostPorts(addr))
}
```

- 在 rpc/item/main.go 中完成初始化操作：
```go
itemServiceImpl := new(ItemServiceImpl)
    // 创建库存客户端
    stockCli, err := NewStockClient("0.0.0.0:8890")
    if err != nil {
       log.Fatal(err)
    }
    // 赋值给商品服务
    itemServiceImpl.stockCli = stockCli

    svr := item.NewServer(itemServiceImpl)
```
使用命令启用商品服务
```sh
sh build.sh
sh output/bootstrap.sh
```
测试接口 打开浏览器访问 localhost:8889/api/item ，结果的 Stock返回了1024，api/main.go中设置的参数