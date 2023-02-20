##  Quick Start
```
安装 goctl 工具
GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@latest
快速生成 api 服务
goctl api new greet
cd greet
go mod init
go mod tidy
go run greet.go -f etc/greet-api.yaml

curl -i http://localhost:8888/from/you
```