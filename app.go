package main

import (
	"embed"
	"html/template"
)

func GenerateMain(FS embed.FS, cfg Config) error {
	t, err := template.ParseFS(FS, "templates/main.go.tmpl")
	if err != nil {
		return err
	}

	outputFile := "cmd/app/main.go"
	f, err := createFile(outputFile)
	if err != nil {
		return err
	}

	if err := t.Execute(f, cfg); err != nil {
		return err
	}

	return f.Close()
}
