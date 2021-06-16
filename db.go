package tuugen

import (
	"html/template"
	"path/filepath"
	"os"
)

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
	for _, dm := range cfg.DataModels {
		t, err := template.ParseFS(FS, "templates/data_model.go.tmpl")
		if err != nil {
			return err
		}

		outputFile := "internal/data/models.go"
		fp := filepath.Dir(outputFile)
		if err := os.MkdirAll(fp, 0777); err != nil {
			return err
		}

		f, err := os.Create(outputFile)
		if err != nil {
			return nil
		}

		if err := t.Execute(f, dm); err != nil {
			return err
		}
	}
}
