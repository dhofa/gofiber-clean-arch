package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

var moduleImport = "github.com/dhofa/gofiber-clean-arch"

type TemplateData struct {
	Name         string
	PascalName   string
	ModuleImport string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go user")
		return
	}
	name := os.Args[1]
	data := TemplateData{
		Name:         name,
		PascalName:   strcase.ToCamel(name),
		ModuleImport: moduleImport,
	}

	files := map[string]string{
		"../../internal/entity/" + name + ".go":                "../../templates/module/entity.tmpl",
		"../../internal/domain/" + name + ".go":                "../../templates/module/domain.tmpl",
		"../../internal/repository/" + name + "_repository.go": "../../templates/module/repository.tmpl",
		"../../internal/usecase/" + name + "_usecase.go":       "../../templates/module/usecase.tmpl",
		"../../internal/handler/" + name + "_handler.go":       "../../templates/module/handler.tmpl",
	}

	for outFile, tmplPath := range files {
		err := generateFile(outFile, tmplPath, data)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Created:", outFile)
		}
	}
}

func generateFile(outputPath, templatePath string, data TemplateData) error {
	tmplBytes, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}

	tmpl, err := template.New("tmpl").Parse(string(tmplBytes))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return err
	}

	// Ensure dir
	os.MkdirAll(strings.TrimSuffix(outputPath, "/"+strings.Split(outputPath, "/")[len(strings.Split(outputPath, "/"))-1]), os.ModePerm)

	return os.WriteFile(outputPath, buf.Bytes(), 0644)
}
