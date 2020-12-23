.PHONY: protos

protos:
	 protoc -I protos/ --go-grpc_out=protos/whatsapp --go_out=protos/whatsapp protos/whatsapp.proto