package main

import (
	"context"
	"log"
	"net"

	"github.com/ahmed82/go-makefile/src/playground/protobuf/proto/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Creat(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte("test"),
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Cannot create Listener: %s", err)
	}

	serverRegisterer := grpc.NewServer()
	service := &myInvoicerServer{}
	//create a type the implement the interface InvoiceServer
	invoicer.RegisterInvoicerServer(serverRegisterer, service)

	err = serverRegisterer.Serve(lis)
	if err != nil {
		log.Fatalf("cannot Serve: %s", err)
	}
}
