clean:
	@echo '$(OS)'
	rm pb/invoicer/*

gen:
	rm pb/*
	protoc --proto_path=proto proto/*.proto  --go_out=:pb --go-grpc_out=:pb --grpc-gateway_out=:pb --openapiv2_out=:swagger

server1:
	go run cmd/server/main.go -port 50051

cert:
	cd cert; ./gen.sh; cd ..

gen_invoicer:
	protoc --go_out=invoicer --go_opt=paths=source_relative  --go-grpc_out=invoicer --go-grpc_opt=paths=source_relative  invoicer.proto

run:
	go run main.go