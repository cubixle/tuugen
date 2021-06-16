package tuugen

type DataModel struct {
	Name       string         `yaml:"name"`
	Properties []DataProperty `yaml:"properties"`
}

type DataProperty struct {
	Name          string `yaml:"name"`
	Type          string `yaml:"type"`
	AutoIncrement bool   `yaml:"autoinc"`
}
