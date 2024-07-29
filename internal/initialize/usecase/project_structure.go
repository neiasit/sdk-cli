package usecase

import (
	"fmt"
	"os"
	"os/exec"
	"sdk-cli/internal/initialize/models"
	"sdk-cli/pkg"
	"sdk-cli/templates"
	"slices"
	"text/template"
)

func CreateProjectStructure(data *models.ProjectData) error {

	dirs := []string{
		"cmd/app",
		"internal",
		"k8s",
		"pkg",
		"migrations",
	}
	for _, dir := range dirs {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	creators := []func(data *models.ProjectData) error{
		createGoMod,
		createMainGo,
	}
	for _, creator := range creators {
		if err := creator(data); err != nil {
			return err
		}
	}

	for _, dir := range dirs {
		exist, err := pkg.CheckIfFilesExist(dir)
		if err != nil {
			return err
		}
		if !exist {
			_, err := os.Create(fmt.Sprintf("%s/%s", dir, ".gitkeep"))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func createGoMod(data *models.ProjectData) error {
	mainGo, err := os.Create("go.mod")
	if err != nil {
		return err
	}

	// Создание шаблона
	templ, err := template.New("main").Funcs(template.FuncMap{
		"ImportPath": ImportPath,
		"HasLibrary": func(lib string) bool {
			return HasLibrary(lib, data.Libraries)
		},
	}).Parse(templates.GoModFileTemplate)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return err
	}
	err = templ.Execute(mainGo, data)
	if err != nil {
		return err
	}

	if err := exec.Command("go", "mod", "tidy").Run(); err != nil {
		return err
	}

	return nil
}

func ImportPath(lib string) string {
	return models.PlatformLibraries[lib]
}

func HasLibrary(lib string, libs []string) bool {
	return slices.Contains(libs, lib)
}

func createMainGo(data *models.ProjectData) error {
	mainGo, err := os.Create("cmd/app/main.go")
	if err != nil {
		return err
	}

	// Создание шаблона
	templ, err := template.New("main").Funcs(template.FuncMap{
		"ImportPath": ImportPath,
		"HasLibrary": func(lib string) bool {
			return HasLibrary(lib, data.Libraries)
		},
	}).Parse(templates.MainFileTemplate)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return err
	}
	err = templ.Execute(mainGo, data)
	if err != nil {
		return err
	}
	return nil
}
