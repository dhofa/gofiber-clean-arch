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
	SnakeName    string
	PascalName   string
	ModuleImport string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go user")
		return
	}

	name := os.Args[1]
	snakeName := strcase.ToSnake(name)
	pascalName := strcase.ToCamel(name)

	data := TemplateData{
		Name:         name,
		SnakeName:    snakeName,
		PascalName:   pascalName,
		ModuleImport: moduleImport,
	}

	basePath := "../../"
	files := map[string]string{
		basePath + "internal/entity/" + snakeName + ".go":                basePath + "templates/module/entity.tmpl",
		basePath + "internal/domain/" + snakeName + ".go":                basePath + "templates/module/domain.tmpl",
		basePath + "internal/repository/" + snakeName + "_repository.go": basePath + "templates/module/repository.tmpl",
		basePath + "internal/usecase/" + snakeName + "_usecase.go":       basePath + "templates/module/usecase.tmpl",
		basePath + "internal/handler/" + snakeName + "_handler.go":       basePath + "templates/module/handler.tmpl",
	}

	// Generate files template
	for outFile, tmplPath := range files {
		err := generateFile(outFile, tmplPath, data)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Created:", outFile)
		}
	}

	// Inject Route in router.go
	err := injectRouteToRouterFile(data.Name, data.Name+"s")
	if err != nil {
		fmt.Println("Error injecting route:", err)
	} else {
		fmt.Println("Route injected to router.go")
	}

	// Inject Handler in router.go
	err = injectHandlerFieldToRegistryStruct(data.Name)
	if err != nil {
		fmt.Println("Error injecting handler:", err)
	} else {
		fmt.Println("Handler injected to router.go")
	}

	// Inject Module Setup in main.go
	err = injectModuleSetupToMainGo(data.Name)
	if err != nil {
		fmt.Println("Error injecting setup:", err)
	} else {
		fmt.Println("Setup injected to main.go")
	}

	// Inject Entity in automigrate.go
	err = injectEntityToAutoMigrate(data.Name)
	if err != nil {
		fmt.Println("Error injecting entity:", err)
	} else {
		fmt.Println("Entity injected to automigrate.go")
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

func injectRouteToRouterFile(handlerName, routeName string) error {
	routeName = strings.ToLower(routeName)
	routeName = strcase.ToKebab(routeName)

	basePath := "../../"
	routerFile := basePath + "infrastructure/router/router.go"
	injection := fmt.Sprintf("\treg.%sHandler.Route(api.Group(\"/%s\"))", strcase.ToCamel(handlerName), routeName)

	input, err := os.ReadFile(routerFile)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")
	var output []string
	inSetup := false
	lastRouteIdx := -1

	for i, line := range lines {
		// Deteksi masuk ke fungsi Setup
		if strings.Contains(line, "func Setup") && strings.Contains(line, "{") {
			inSetup = true
		}

		// Temukan baris terakhir reg.XHandler.Route(...)
		if inSetup && strings.Contains(line, ".Route(") {
			lastRouteIdx = i
		}

		// Deteksi keluar dari Setup
		if inSetup && strings.TrimSpace(line) == "}" {
			inSetup = false
		}

		output = append(output, line)
	}

	// Jika ditemukan baris .Route(...) sebelumnya, inject setelahnya
	if lastRouteIdx != -1 {
		// Sisipkan di index setelah lastRouteIdx
		before := output[:lastRouteIdx+1]
		after := output[lastRouteIdx+1:]
		output = append(before, append([]string{injection}, after...)...)
	}

	return os.WriteFile(routerFile, []byte(strings.Join(output, "\n")), 0644)
}

func injectHandlerFieldToRegistryStruct(module string) error {
	basePath := "../../"
	filePath := basePath + "infrastructure/router/router.go"
	handlerField := fmt.Sprintf("\t%sHandler *handler.%sHandler", strcase.ToCamel(module), strcase.ToCamel(module))

	input, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")
	var output []string
	inStruct := false
	inserted := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Deteksi awal struct RouteRegistry
		if strings.HasPrefix(trimmed, "type RouteRegistry struct") {
			inStruct = true
		}

		// Jika dalam struct dan menemukan penutup, sisipkan sebelum "}"
		if inStruct && trimmed == "}" && !inserted {
			output = append(output, handlerField)
			inserted = true
		}

		output = append(output, line)

		// Reset flag setelah keluar struct
		if inStruct && trimmed == "}" {
			inStruct = false
		}
	}

	return os.WriteFile(filePath, []byte(strings.Join(output, "\n")), 0644)
}

func injectModuleSetupToMainGo(module string) error {
	basePath := "../../"
	filePath := basePath + "cmd/main.go"
	pascal := strcase.ToCamel(module)
	camel := strcase.ToLowerCamel(module)

	// Baris kode setup module
	setupLines := []string{
		fmt.Sprintf("\t// %s setup", pascal),
		fmt.Sprintf("\t%sRepo := repository.New%sRepository(database)", camel, pascal),
		fmt.Sprintf("\t%sUsecase := usecase.New%sUsecase(%sRepo)", camel, pascal, camel),
		fmt.Sprintf("\t%sHandler := handler.New%sHandler(%sUsecase)\n", camel, pascal, camel),
	}

	// Baris untuk dimasukkan ke dalam RouteRegistry struct
	assignLine := fmt.Sprintf("\t\t%sHandler: %sHandler,", pascal, camel)

	input, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")
	var output []string
	insertedSetup := false
	inRouteAssign := false
	closingBraceIdx := -1

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Inject setup block sebelum komentar: "// Inject all handlers ke router registry"
		if !insertedSetup && strings.Contains(trimmed, "// Inject all handlers") {
			output = append(output, setupLines...)
			insertedSetup = true
		}

		// Deteksi awal blok RouteRegistry
		if strings.Contains(trimmed, "routes := &router.RouteRegistry{") {
			inRouteAssign = true
		}

		// Tandai posisi penutup blok RouteRegistry
		if inRouteAssign && trimmed == "}" {
			closingBraceIdx = len(output)
			inRouteAssign = false
		}

		output = append(output, line)
	}

	// Sisipkan assignLine sebelum penutup blok
	if closingBraceIdx != -1 {
		output = append(output[:closingBraceIdx], append([]string{assignLine}, output[closingBraceIdx:]...)...)
	}

	return os.WriteFile(filePath, []byte(strings.Join(output, "\n")), 0644)
}

func injectEntityToAutoMigrate(module string) error {
	basePath := "../../"
	filePath := basePath + "infrastructure/db/postgres.go"
	entityLine := fmt.Sprintf("\tdb.AutoMigrate(&entity.%s{})", strcase.ToCamel(module))

	input, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")
	var output []string
	lastAutoMigrateIdx := -1

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		// Deteksi baris db.AutoMigrate(...)
		if strings.HasPrefix(trimmed, "db.AutoMigrate(&entity.") {
			lastAutoMigrateIdx = i
		}
		output = append(output, line)
	}

	// Jika ada baris AutoMigrate ditemukan, sisipkan tepat di bawahnya
	if lastAutoMigrateIdx != -1 {
		before := output[:lastAutoMigrateIdx+1]
		after := output[lastAutoMigrateIdx+1:]
		output = append(before, append([]string{entityLine}, after...)...)
	}

	return os.WriteFile(filePath, []byte(strings.Join(output, "\n")), 0644)
}
