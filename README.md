`cmd/api/main.go`启动API service；

`cmd/rpc/main.go`启动RPC serivce。

接口定义位于`api/kitex/oasyce.thrift`，例如`/GetUser`可以构造为`http://localhost:8080/?Action=GetUser`