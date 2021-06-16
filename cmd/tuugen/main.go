package main

import (
	"log"

	"github.com/cubixle/tuugen"
)

func main() {
	// idea is to read a yaml file and create the project.
	err := tuugen.GenerateProtos()
	if err != nil {
		log.Fatal(err)
	}
	err = tuugen.GenerateService("./internal/pb/service_grpc.pb.go")
	if err != nil {
		log.Fatal(err)
	}
}
