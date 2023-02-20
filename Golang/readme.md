# Golang

## History
```
从 Go 1.5 版本的自举（即用 Go 语言编写程序来实现 Go 语言自身），
到 Go 1.7 版本的极速 GC（也称垃圾回收器），
到 2018 年 2 月发布的 Go 1.10 版本对其自带工具的全面升级，
以及可预见的后续版本关键特性（比如用来做程序依赖管理的go mod命令
```
## Prepare
- tour: https://tour.go-zh.org/welcome/1
- geek time: https://github.com/hyper0x/Golang_Puzzlers
- go_command_tutorial: https://github.com/hyper0x/go_command_tutorial
## Debug
```
brew install delve
dlv version
```
- debug with Gin
dlv debug main.go
b code/path

## Gen Code of Model 
- https://github.com/smallnest/gen
```
go install github.com/smallnest/gen@latest     

~/go/bin/gen --sqltype=postgres \
   	--connstr "postgresql://nomad_logistics_dev:" \
   	--database main  \
    --table fc_configuration_template \
   	--json \
   	--gorm \
   	--guregu \
   	--rest \
   	--out ./example \
   	--module example.com/rest/example \
   	--mod \
   	--server \
   	--makefile \
   	--json-fmt=snake \
   	--generate-dao \
   	--generate-proj \
   	--overwrite
```