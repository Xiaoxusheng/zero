package zero

//api文件格式化，全局通用
//go:generate goctl api format --dir ./user/api/user.api

/*----------------------------------------------------------------------------------*/
//生成user.api文件
//go:generate goctl api  -o ./user/api/user.api

//生成api  user
//go:generate goctl api go -api ./user/api/user.api -dir=./user/api

//生成user.proto
//go:generate goctl rpc  --o ./user/rpc/user.proto

//生成user.rpc
//go:generate goctl rpc protoc ./user/rpc/user.proto --go_out=./user/rpc --go-grpc_out=./user/rpc --zrpc_out=./user/rpc

//goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.

//运行 api
//go:generate   go run user.go -f etc/user-api.yaml
//运行rpc
//go:generate  go run user.go -f etc/user.yaml

//-------------------------------------------------------------
//生成video.api文件
//go:generate goctl api  -o ./video/api/video.api

//生成api
//go:generate goctl api go -api ./video/api/video.api -dir=./video/api

//生成video。proto文件
//go:generate goctl rpc  --o ./video/rpc/video.proto

//生成video.rpc
//go:generate goctl rpc protoc ./video/rpc/video.proto --go_out=./video/rpc --go-grpc_out=./video/rpc --zrpc_out=./video/rpc
