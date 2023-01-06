proto:
	protoc pkg/**/pb/*.proto --go-grpc_out=. --go_out=.

gateway:
	go run gateway/main.go

user-server:
	go run pkg/user/cmd/main.go