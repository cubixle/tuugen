package tuugen

import (
	"embed"
	"html/template"
	"os"
	"path/filepath"
	"strings"
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

func (p DataProperty) NameCorrected() string {
	if strings.Contains(p.Name, "_") {
		parts := strings.Split(p.Name, "_")
		newParts := []string{}
		for _, p := range parts {
			newParts = append(newParts, strings.Title(p))
		}
		return strings.Join(newParts, "")
	}
	return strings.Title(p.Name)
}

func (p DataProperty) DBTypeToGoType() string {
	switch p.Type {
	case "varchar":
		return "string"
	case "timestamp":
		return "time.Time"
	case "int":
		return "int32"
	default:
		return "string"
	}
}

func GenerateDataModels(FS embed.FS, cfg Config) error {
	for _, dm := range cfg.DataModels {
		t, err := template.ParseFS(FS, "templates/data_model.go.tmpl")
		if err != nil {
			return err
		}

		outputFile := "internal/storage/" + strings.ToLower(dm.Name) + ".go"
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
	return nil
}
