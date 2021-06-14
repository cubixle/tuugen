package tuugen

func GenerateProtos() error {
	cmd := "protoc"
	args := []string{"-I", ".:${GOPATH}/src", "--go_out=.", "--go-grpc_out=.", "service.proto"}
	return runCommand(cmd, args)
}
