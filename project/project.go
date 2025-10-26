package project

import (
	"donbarrigon/make/cmd"
	"os"
	"regexp"
	"strings"
)

func New() {
	projectUser, projectName := getProjectName()
	projectType, projectOptions := getProjectType()

	switch projectType {
	case "Go":
		createGoProject(projectUser, projectName, projectOptions)
	default:
		cmd.Danger("TODO: tipo de proyecto " + projectType + " No implementado")
	}

}

func getProjectName() (user string, name string) {
	gitRepoPattern := regexp.MustCompile(`^[a-zA-Z0-9_-]+/[a-zA-Z0-9_-]+$`)
	cmd.Print("Formato <usuario-de-git>/<nombre-del-proyecto>")
	project := cmd.Prompt("Nombre del proyecto", "")
	if !gitRepoPattern.MatchString(project) {
		cmd.Danger("El nombre del proyecto debe ser en formato <usuario-de-git>/<nombre-del-proyecto>")
		os.Exit(2)
	}
	p := strings.Split(project, "/")
	projectUser := p[0]
	projectName := p[1]

	return projectUser, projectName
}

func getProjectType() (string, []string) {
	cmd.Info("Configuración del proyecto")

	tipe := cmd.PromptSelect("🏗️ Selecciona el tipo de proyecto", 0, "Go", "Vue", "Bun")
	options := []string{}
	if tipe == "Go" {
		if cmd.PromptConfirm("Fork al proyecto original?", false) {
			options = append(options, "fork")
		}
		if cmd.PromptConfirm("Usar GORM?", false) {
			options = append(options, "gorm")
		}
		if cmd.PromptConfirm("Instalar Documentacion?", true) {
			options = append(options, "docs")
		}
	}
	if cmd.PromptConfirm("Usar TypeScript?", false) {
		options = append(options, "ts")
		cmd.Warning("Siempre hay gente dispuesta a trabajar en modo dificil.")
	} else {
		options = append(options, "js")
	}
	if cmd.PromptConfirm("Instalar tailwind?", false) {
		options = append(options, "tailwind")
	}

	return tipe, options
}
