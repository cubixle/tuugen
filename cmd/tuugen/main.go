package main

import (
	"embed"
	"io/ioutil"
	"log"

	"github.com/cubixle/tuugen"
)

//go:embed templates/*
var FS embed.FS

func main() {
	// read config file
	d, err := ioutil.ReadFile("tuugen.yml")
	if err != nil {
		log.Fatalf("failed to find tuugen.yaml, %v", err)
	}
	cfg, err := tuugen.YamlToConfig(d)
	if err != nil {
		log.Fatalf("failed to load config, %v", err)
	}

	log.Println("Generating protos")
	err = tuugen.GenerateProtos()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Generating service implementation")
	err = tuugen.GenerateService(FS, cfg)
	if err != nil {
		log.Fatalf("failed to generate service: %v", err)
	}
	log.Println("Generating interactor stubs")
	err = tuugen.GenerateInteractor(FS, cfg)
	if err != nil {
		log.Fatalf("failed to generate interactor: %v", err)
	}

	log.Println("Generating data models")
	if err := tuugen.GenerateDataModels(FS, cfg); err != nil {
		log.Fatalf("failed to generate data models: %v", err)
	}

	log.Printf("Running 'go mod init %s' \n", cfg.ImportPath)
	if err := tuugen.GoModInit(cfg.ImportPath); err != nil {
		log.Fatalf("failed to run 'go mod init': %v", err)
	}
	log.Println("Running clean up")
	if err := tuugen.GoFmt(); err != nil {
		log.Fatalf("failed to run gofmt: %v", err)
	}
	if err := tuugen.GoImports(); err != nil {
		log.Fatalf("failed to run 'goimports': %v", err)
	}
	log.Println("------")
	log.Println("🎉 All setup has been complete enjoy working on your business logic.")
	log.Println("------")
}
