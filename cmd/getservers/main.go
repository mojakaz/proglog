package main

import (
	"context"
	"flag"
	"fmt"
	api "github.com/mojakaz/proglog/api/v1"
	"github.com/mojakaz/proglog/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	addr := flag.String("addr", ":8400", "service address")
	flag.Parse()
	peerTLSConfig, err := config.SetupTLSConfig(config.TLSConfig{
		CertFile:      config.RootClientCertFile,
		KeyFile:       config.RootClientKeyFile,
		CAFile:        config.CAFile,
		Server:        false,
		ServerAddress: "127.0.0.1",
	})
	tlsCreds := credentials.NewTLS(peerTLSConfig)
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(tlsCreds))
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewLogClient(conn)
	ctx := context.Background()
	res, err := client.GetServers(ctx, &api.GetServersRequest{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("servers:")
	for _, server := range res.Servers {
		fmt.Printf("\t- %v\n", server)
	}
}
