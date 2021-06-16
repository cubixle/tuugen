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
	// idea is to read a yaml file and create the project.
	err = tuugen.GenerateProtos()
	if err != nil {
		log.Fatal(err)
	}

	err = tuugen.GenerateService(FS, cfg)
	if err != nil {
		log.Fatalf("failed to generate service: %v", err)
	}
	err = tuugen.GenerateInteractor(FS, cfg)
	if err != nil {
		log.Fatalf("failed to generate interactor: %v", err)
	}
	if err := tuugen.GoFmt(); err != nil {
		log.Fatalf("failed to run gofmt: %v", err)
	}
	if err := tuugen.GoModInit(cfg.ImportPath); err != nil {
		log.Fatalf("failed to run 'go mod init': %v", err)
	}
	if err := tuugen.GoImports(); err != nil {
		log.Fatalf("failed to run 'goimports': %v", err)
	}
}
