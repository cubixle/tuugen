package main

import (
	"embed"
	_ "embed"
	"log"

	"github.com/cubixle/tuugen"
)

//go:embed templates/*
var FS embed.FS

func main() {

	// idea is to read a yaml file and create the project.
	err := tuugen.GenerateProtos()
	if err != nil {
		log.Fatal(err)
	}
	importPath := "github.com/cubixle/tuugen/example_project"

	err = tuugen.GenerateService(
		FS,
		"Service",
		importPath,
		"./internal/pb/service/service_grpc.pb.go",
		"./internal/service/service.go",
	)
	if err != nil {
		log.Fatalf("failed to generate service: %v", err)
	}
	err = tuugen.GenerateInteractor(
		FS,
		"Service",
		importPath,
		"./internal/pb/service/service_grpc.pb.go",
		"./internal/interactors/interactors.go",
	)
	if err != nil {
		log.Fatalf("failed to generate interactor: %v", err)
	}

	if err := tuugen.GoFmt(); err != nil {
		log.Fatalf("failed to run gofmt: %v", err)
	}
	if err := tuugen.GoModInit(importPath); err != nil {
		log.Fatalf("failed to run 'go mod init': %v", err)
	}
	if err := tuugen.GoImports(); err != nil {
		log.Fatalf("failed to run 'goimports': %v", err)
	}
}
