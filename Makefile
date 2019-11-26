PROTOC3 :=$(shell which protoc)

all: proto client server gateway

.PHONY: client
client:
	@CGO_ENABLED=0 go build -ldflags="-s -w" -o build/client ./greeter_client/main.go	

.PHONY: server
server:
	@CGO_ENABLED=0 go build -ldflags="-s -w" -o build/server ./greeter_server/main.go	

.PHONY: gateway
gateway:
	@CGO_ENABLED=0 go build -ldflags="-s -w" -o build/gateway ./gateway/main.go	

proto:
	$(PROTOC3) -I/usr/local/include -I. \
		-I$$GOPATH/src \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
		--go_out=plugins=grpc:.  \
		--swagger_out=logtostderr=true:. \
		--govalidators_out=. \
		--grpc-gateway_out=logtostderr=true:. \
		proto/*.proto


.PHONY: clean
clean:
	@rm -rf build/*
