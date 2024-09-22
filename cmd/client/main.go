package main

import (
	"flag"
	"log"

	"github.com/Roayuso/gRCP_APIS/client"
	"github.com/Roayuso/gRCP_APIS/sample"

	"google.golang.org/grpc"
)

func testCreateLaptop(laptopClient *client.LaptopClient) {
	laptopClient.CreateLaptop(sample.NewLaptop())
}

func testUploadImage(laptopClient *client.LaptopClient) {
	laptop := sample.NewLaptop()
	laptopClient.CreateLaptop(laptop)
	laptopClient.UploadImage(laptop.GetId(), "tmp/laptop.jpg")
}

func main() {
	serverAddress := flag.String("address", "", "the server address")
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	flag.Parse()
	log.Printf("dial server %s, TLS = %t", *serverAddress, *enableTLS)

	transportOption := grpc.WithInsecure()

	cc2, err := grpc.Dial(
		*serverAddress,
		transportOption,
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := client.NewLaptopClient(cc2)
	testCreateLaptop(laptopClient)
	testUploadImage(laptopClient)

}
