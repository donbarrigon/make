package project

import (
	"donbarrigon/make/cmd"
	"donbarrigon/make/config"
	"os"
	"regexp"
	"strings"
)

func New() {
	projectUser, projectName := getProjectName()
	projectType, projectOptions := getProjectType()

	switch projectType {
	case "go":
		createGoProject(projectUser, projectName, projectOptions)
	default:
		cmd.Danger("TODO: No implementado")
	}

	config.Save(map[string]any{
		"user":    projectUser,
		"name":    projectName,
		"type":    projectType,
		"options": projectOptions,
	})

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

	tipe := cmd.PromptSelect("🏗️ Selecciona el tipo de proyecto", 0, "go", "vue", "bun")
	options := []string{}
	if tipe == "go" {
		if cmd.PromptConfirm("Fork al proyecto original?", false) {
			options = append(options, "fork")
		}
		if cmd.PromptConfirm("Instalar Documentacion?", true) {
			options = append(options, "docs")
			cmd.Warning("requiere que instales la dependencia alpine")
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
	if cmd.PromptConfirm("Instalar Alpine?", true) {
		options = append(options, "alpine")
	}

	return tipe, options
}
