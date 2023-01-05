package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/victorgcramos/zi/pkg/config"
	"github.com/victorgcramos/zi/pkg/user"
	"github.com/victorgcramos/zi/pkg/user/pb"
	"google.golang.org/grpc"
)

func _main() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}
	// Setup net listener
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalln("[USER]: Failed to listing:", err)
	}
	fmt.Println("[USER]: User service on", cfg.Port)
	//  Setup GRPC Server
	grpcs := grpc.NewServer()
	// 	Setup User Service
	us := user.NewService()
	pb.RegisterUserServiceServer(grpcs, us)

	//  Setup TCP Server
	if err := grpcs.Serve(lis); err != nil {
		return err
	}
	return nil
}

func main() {
	err := _main()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
