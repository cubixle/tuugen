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

func GenerateDataModels(cfg Config) error {
	t, err := template.ParseFS(FS, "templates/data_model.go.tmpl")
	if err != nil {
		return err
	}

	outputFile := "internal/data/models.go"
	// check/create file path
	fp := filepath.Dir(outputFile)
}