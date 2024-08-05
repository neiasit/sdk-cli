package ui

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/labstack/gommon/color"
	"github.com/neiasit/sdk-cli/internal/initialize/models"
	"github.com/neiasit/sdk-cli/internal/initialize/usecase"
	"os"
	"strings"
)

func formatPlatformLibs(platformLibs map[string]string) []string {
	var titles []string
	for t := range platformLibs {
		titles = append(titles, t)
	}
	return titles
}

func DisplayMenu(this bool) error {
	var options models.ProjectData

	questions := []*survey.Question{
		{
			Name:     "projectName",
			Prompt:   &survey.Input{Message: "Enter project name:"},
			Validate: survey.Required,
		},
		{
			Name: "golangVersion",
			Prompt: &survey.Select{
				Message: "Choose golang version:",
				Options: models.GolangVersions,
			},
		},
		{
			Name: "libraries",
			Prompt: &survey.MultiSelect{
				Message: "Choose platform libraries:",
				Options: formatPlatformLibs(models.PlatformLibraries),
			},
		},
		{
			Name: "otherOptions",
			Prompt: &survey.MultiSelect{
				Message: "Choose other options:",
				Options: models.AdditionalOptions,
			},
		},
	}

	err := survey.Ask(questions, &options)
	if err != nil {
		return err
	}

	options.ProjectName = strings.ToLower(options.ProjectName)

	println(color.Green(
		"New project settings:\n"+fmt.Sprintf(
			"Project name: %s\n"+
				"Golang Version: %s\n"+
				"Platfrom Libraries: %s\n"+
				"Other Options: %s\n",
			options.ProjectName,
			options.GolangVersion,
			options.Libraries,
			options.OtherOptions,
		), color.B))

	var isCorrect bool
	if err := survey.AskOne(&survey.Confirm{
		Message: "Are these settings correct?",
	}, &isCorrect); err != nil {
		return err
	}
	if !isCorrect {
		return DisplayMenu(this)
	}

	if !this {
		err = os.Mkdir(options.ProjectName, os.ModePerm)
		if err != nil {
			return err
		}
		err = os.Chdir(options.ProjectName)
		if err != nil {
			return err
		}
	}

	if err := usecase.CreateProjectStructure(&options); err != nil {
		return err
	}
	if err := usecase.CreateAdditionalOptions(&options); err != nil {
		return err
	}
	println(color.Green("Project successful created", color.B))
	return nil
}
