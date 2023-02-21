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
### Windows
Download the win64.zip file from the release 
https://github.com/protocolbuffers/protobuf/releases/tag/v22.0
and the new Path variable point to the bin folder

Test in the CMD:
```
protoc --version
```
image.png

- Install `protoc-gen-go` and `protoc-gen-go-grpc`
* Go plugins for the protocol compiler
```bash
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

protobuf devSite 
https://protobuf.dev/getting-started/gotutorial/
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
```
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
```
- Install `protoc-gen-grpc-gateway` and `protoc-gen-openapiv2`

```bash
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
```

using the GitBash works:
```bash
protoc --go_out=. --go_opt=paths=source_relative    --go-grpc_out=. --go-grpc_opt=paths=source_relative  invoicer.proto

protoc --go_out=invoicer --go_opt=paths=source_relative  --go-grpc_out=invoicer --go-grpc_opt=paths=source_relative  invoicer.proto
```

download the missing lib
```
go get -u google.golang.org/grpc
go mod tidy
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