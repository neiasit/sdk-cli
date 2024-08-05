package usecase

import (
	"fmt"
	"github.com/neiasit/sdk-cli/internal/initialize/models"
	"github.com/neiasit/sdk-cli/templates"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

func CreateAdditionalOptions(data *models.ProjectData) error {
	data.ProjectName = strings.ToLower(data.ProjectName)
	for _, option := range data.OtherOptions {
		switch option {
		case models.GitInitializationOption:
			if err := initializeGit(data); err != nil {
				return err
			}
		case models.DockerOption:
			if err := initializeDocker(data); err != nil {
				return err
			}
		case models.GithubCiCdOption:
			if err := initializeGithubCiCd(data); err != nil {
				return err
			}
		}
	}

	return nil
}

func initializeGithubCiCd(data *models.ProjectData) error {

	if err := os.MkdirAll(".github/workflows", os.ModePerm); err != nil {
		return err
	}

	mainGo, err := os.Create(".github/workflows/go.yml")
	if err != nil {
		return err
	}

	// Создание шаблона
	templ, err := template.New("github-workflow").Funcs(template.FuncMap{
		"ImportPath": ImportPath,
		"HasLibrary": func(lib string) bool {
			return HasLibrary(lib, data.Libraries)
		},
	}).Parse(templates.GithubCiCdWorkflowTemplate)
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

func initializeDocker(data *models.ProjectData) error {
	mainGo, err := os.Create("Dockerfile")
	if err != nil {
		return err
	}

	// Создание шаблона
	templ, err := template.New("dockerfile").Funcs(template.FuncMap{
		"ImportPath": ImportPath,
		"HasLibrary": func(lib string) bool {
			return HasLibrary(lib, data.Libraries)
		},
	}).Parse(templates.DockerfileTemplate)
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

func initializeGit(data *models.ProjectData) error {
	if err := exec.Command("git", "init").Run(); err != nil {
		return err
	}
	mainGo, err := os.Create(".gitignore")
	if err != nil {
		return err
	}

	// Создание шаблона
	templ, err := template.New("main").Funcs(template.FuncMap{
		"ImportPath": ImportPath,
		"HasLibrary": func(lib string) bool {
			return HasLibrary(lib, data.Libraries)
		},
	}).Parse(templates.GitignoreTemplate)
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
