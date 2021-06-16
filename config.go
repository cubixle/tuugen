package tuugen

import "gopkg.in/yaml.v2"

type Config struct {
	Project     string      `yaml:"project"`
	ServiceName string      `yaml:"service_name"`
	ImportPath  string      `yaml:"import_path"`
	ProtoFile   string      `yaml:"proto_file"`
	DataModels  []DataModel `yaml:"data_models"`
}

func YamlToConfig(b []byte) (Config, error) {
	cfg := Config{}
	err := yaml.Unmarshal(b, &cfg)
	return cfg, err
}
