package main

import "os/exec"

func runCommand(cmd string, args []string) error {
	c := exec.Command(cmd, args...)
	err := c.Start()
	if err != nil {
		return nil
	}
	err = c.Wait()
	if err != nil {
		return err
	}
	return nil
}

func GenerateProtos(protoFile string) error {
	cmd := "protoc"
	args := []string{"-I", ".:${GOPATH}/src", "--go_out=.", "--go-grpc_out=.", protoFile}
	return runCommand(cmd, args)
}

func GoFmt() error {
	cmd := "gofmt"
	return runCommand(cmd, []string{})
}

func GoImports() error {
	if err := runCommand("goimports", []string{"-w", "internal/service/service.go"}); err != nil {
		return err
	}
	if err := runCommand("goimports", []string{"-w", "internal/storage/"}); err != nil {
		return err
	}
	return runCommand("goimports", []string{"-w", "internal/interactors/interactors.go"})
}

func GoModInit(importPath string) error {
	runCommand("go", []string{"mod", "init", importPath})
	return runCommand("go", []string{"mod", "tidy"})
}
