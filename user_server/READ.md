```bash
# 生成rpc代码
goctl rpc protoc ./api/user_server.proto --go_out=. --go-grpc_out=. --zrpc_out=. --client=true --style=goZero  --multiple
```

```bash
# 运行
go run ./cmd/userserver.go
```

```bash
# 编译 服务端
go build -o ./build/UserServer.exe ./cmd/userserver.go
```
```bash
# 运行服务端
./build/UserServerBuild.exe
```

```bash
# 编译 迁移工具
go build -o ./build/UserModel.exe ./migration/UserDatabaseMigration.go
```
```bash
# 运行迁移工具
./build/UserModel.exe
```
```bash
# 构建镜像
docker build -t go_user_server:v1 .
```