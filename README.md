# 环境
## go版本
   - >1.5

# 安装

## protobuf 3以上
```
   $git clone https://github.com/google/protobuf
   $cd protobuf
   $./autogen.sh
   $./configure
   $make
   $make check
   #make install
```

## go grpc相关包

```
   go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
   go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
   go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
   go get google.golang.org/grpc
```

# proto目录

## google api
  - 问题protoc-gen-grpc-gateway中第三方的google api报错

## 自定义
   helloworld.proto

## 生成grpc和grpc-gateway的go文件
   make 
   

# 测试

## 启动greeter server
```
   $cd greeter_server && go run main.go
```

## 启动gateway
```
   $cd gateway && go run main.go
```

## 模拟http请求
```
   $curl -X POST -k http://localhost:8080/v1/example/echo -d '{"name": " world"}'
```

# 仅供参考
   - [项目](https://www.cnblogs.com/andyidea/p/6529900.html)
   - [带https](https://github.com/EDDYCJY/grpc-hello-world)
