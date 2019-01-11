
PROTOC3 ?= /opt/protobuf/bin/protoc

PROTO_FILE = helloworld.proto

DIR=$(shell pwd)
#DIR = grpc-helloworld-gateway/proto

all:

	$(PROTOC3) -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto
	#编译hello_http.proto为hello_http.pb.proto
	$(PROTOC3) -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=$(DIR)/google/api:. $(PROTO_FILE)
	#编译hello_http.proto为hello_http.pb.gw.proto
	$(PROTOC3) --grpc-gateway_out=logtostderr=true:. $(PROTO_FILE)

