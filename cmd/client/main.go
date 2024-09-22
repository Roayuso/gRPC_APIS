package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Roayuso/gRPC_APIS/client"
	"github.com/Roayuso/gRPC_APIS/sample"

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

func testGetLaptopStore(laptopClient *client.LaptopClient) {
	laptops := laptopClient.GetLaptopStore()
	for _, laptop := range laptops {
		fmt.Println(laptop.Cpu.Name)
	}
}

func main() {
	serverAddress := "0.0.0.0:52454"
	flag.Parse()
	log.Printf("dial server %s", serverAddress)

	transportOption := grpc.WithInsecure()

	cc2, err := grpc.Dial(
		serverAddress,
		transportOption,
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := client.NewLaptopClient(cc2)
	testCreateLaptop(laptopClient)
	// testUploadImage(laptopClient)
	testGetLaptopStore(laptopClient)

}
