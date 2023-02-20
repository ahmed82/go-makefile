clean:
	rm pb/*
	rm swagger/*

gen:
	protoc --proto_path=proto proto/*.proto  --go_out=:pb --go-grpc_out=:pb --grpc-gateway_out=:pb --openapiv2_out=:swagger

server1:
	go run cmd/server/main.go -port 50051