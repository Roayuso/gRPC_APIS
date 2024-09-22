package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/Roayuso/gRPC_APIS/pb/proto"
	"github.com/Roayuso/gRPC_APIS/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func runGRPCServer(
	laptopServer pb.LaptopServiceServer,
	listener net.Listener,
) error {

	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	reflection.Register(grpcServer)

	log.Printf("Start GRPC server at %s", listener.Addr().String())
	return grpcServer.Serve(listener)
}

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")

	laptopServer := service.NewLaptopServer(laptopStore, imageStore)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = runGRPCServer(laptopServer, listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
