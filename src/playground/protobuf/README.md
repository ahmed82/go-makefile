# install protobuf

```
brew install protobuf
```

## navigate intop `https://grpc.io/`


## Setup development environment

- Install `protoc`:

```bash
brew install protobuf
```

- Install `protoc-gen-go` and `protoc-gen-go-grpc`

```bash
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

- Install `protoc-gen-grpc-gateway` and `protoc-gen-openapiv2`

```bash
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
```

## How to run

- Generate codes from proto files:

```bash
make gen
```

- Clean up generated codes in `pb` and `swagger` folders:

```bash
make clean
```
https://github.com/techschool/pcbook-go