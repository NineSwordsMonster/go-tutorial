# Build

````shell
 go build
````

## 发现Go的安装路径，go命令将在那里安装当前的软件包
```shell
# 
go list -f '{{.Target}}'
go env -w GOBIN=/path/to/your/bin
 go install 
 hello
```