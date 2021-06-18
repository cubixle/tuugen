package main

import (
    "os"
	"fmt"
	"net"
	"log"
	
    "google.golang.org/grpc"
	pb "github.com/cubixle/tuugen/example_project/internal/pb/service"
	"github.com/cubixle/tuugen/example_project/internal/service"
)

func main() {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%s", getEnv("GRPC_PORT", "8090")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceServer(s, service.New(store))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getEnv(name, defaultValue string) string {
    v := os.Getenv(name)
    if v == "" {
        return defaultValue
    }
    return v
}