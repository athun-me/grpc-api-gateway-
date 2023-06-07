export GO111MODULE=on

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/auth/pb/auth.proto

run:
	go run cmd/main.go