package main

import (
	"embed"
	"io/ioutil"
	"log"
	"time"
)

//go:embed templates/*
var FS embed.FS

func main() {
	tNow := time.Now()
	// read config file
	d, err := ioutil.ReadFile("tuugen.yml")
	if err != nil {
		log.Fatalf("failed to find tuugen.yaml, %v", err)
	}
	cfg, err := YamlToConfig(d)
	if err != nil {
		log.Fatalf("failed to load config, %v", err)
	}

	log.Printf("Generating protos from %s\n", cfg.ProtoFile)
	err = GenerateProtos(cfg.ProtoFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Generating service implementation")
	err = GenerateService(FS, cfg)
	if err != nil {
		log.Fatalf("failed to generate service: %v", err)
	}
	log.Println("Generating interactor stubs")
	err = GenerateInteractor(FS, cfg)
	if err != nil {
		log.Fatalf("failed to generate interactor: %v", err)
	}

	log.Println("Generating data models")
	if err := GenerateDataModels(FS, cfg); err != nil {
		log.Fatalf("failed to generate data models: %v", err)
	}

	log.Println("Generate main.go")
	if err := GenerateMain(FS, cfg); err != nil {
		log.Fatalf("failed to generate main file: %v", err)
	}

	log.Printf("Running 'go mod init %s' \n", cfg.ImportPath)
	if err := GoModInit(cfg.ImportPath); err != nil {
		log.Fatalf("failed to run 'go mod init': %v", err)
	}
	log.Println("Running clean up")
	if err := GoFmt(); err != nil {
		log.Fatalf("failed to run gofmt: %v", err)
	}
	if err := GoImports(); err != nil {
		log.Fatalf("failed to run 'goimports': %v", err)
	}
	log.Println("------")
	log.Printf("🎉 All setup has been complete enjoy working on your business logic. took: %s\n", time.Since(tNow))
	log.Println("------")
}
