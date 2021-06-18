package main

import (
	"embed"
	"html/template"
	"os"
	"path/filepath"
)

func GenerateMain(FS embed.FS, cfg Config) error {
	t, err := template.ParseFS(FS, "templates/main.go.tmpl")
	if err != nil {
		return err
	}

	outputFile := "cmd/app/main.go"
	fp := filepath.Dir(outputFile)
	if err := os.MkdirAll(fp, 0777); err != nil {
		return err
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return nil
	}

	if err := t.Execute(f, cfg); err != nil {
		return err
	}
	return nil
}
