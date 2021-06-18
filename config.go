package main

import (
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Project     string      `yaml:"project"`
	ServiceName string      `yaml:"service_name"`
	ImportPath  string      `yaml:"import_path"`
	ProtoFile   string      `yaml:"proto_file"`
	GRPCFile    string      `yaml:"grpc_file"`
	DataModels  []DataModel `yaml:"data_models"`
}

func (c Config) GRPCPath() string {
	return filepath.Dir(c.GRPCFile)
}

func YamlToConfig(b []byte) (Config, error) {
	cfg := Config{}
	err := yaml.Unmarshal(b, &cfg)
	return cfg, err
}
