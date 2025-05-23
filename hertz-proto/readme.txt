https://blog.csdn.net/chenhuyu7/article/details/146139569?ops_request_misc=&request_id=&biz_id=102&utm_term=go%E4%BD%BF%E7%94%A8proto&utm_medium=distribute.pc_search_result.none-task-blog-2~all~sobaiduweb~default-0-146139569.142^v102^pc_search_result_base7&spm=1018.2226.3001.4187

准备 ./idl 文件夹下的文件

下载操作系统下的 protof
[window](https://github.com/protocolbuffers/protobuf/releases/download/v30.0/protoc-30.0-win64.zip)

protoc --go_out=. --go-grpc_out=. idl/api.proto
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go-grpc@latest
// 安装的程序 可以通过 go env GOPATH 查看

hz new -module example.com/m -I idl -idl idl/hello/hello.proto
go mod tidy