package main

import (
	"embed"
	"html/template"
	"strings"
)

type DataModel struct {
	Name       string         `yaml:"name"`
	Properties []DataProperty `yaml:"properties"`
}

type DataProperty struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

func (p DataProperty) NameCorrected() string {
	if !strings.Contains(p.Name, "_") {
		return strings.Title(p.Name)
	}

	parts := strings.Split(p.Name, "_")
	newParts := []string{}
	for _, p := range parts {
		newParts = append(newParts, strings.Title(p))
	}
	return strings.Join(newParts, "")
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
		f, err := createFile(outputFile)
		if err != nil {
			return err
		}
		defer f.Close()

		if err := t.Execute(f, dm); err != nil {
			return err
		}
	}
	return nil
}
