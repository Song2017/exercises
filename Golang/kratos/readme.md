## guide
### quick-start
https://go-zero.dev/cn/docs/quick-start/monolithic-service
### go install 安装：
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
kratos upgrade


### Create a demo
```
# 创建项目模板
kratos new helloworld

cd helloworld
# 拉取项目依赖
go mod tidy

# 运行程序
kratos run
```