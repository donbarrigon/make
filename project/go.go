package project

import (
	"donbarrigon/make/cmd"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func createGoProject(projectUser string, projectName string, options []string) {
	cmd.Runf("Clonando el repositorio", "git", "clone", "https://github.com/donbarrigon/new.git "+projectName)
	cmd.Run("cd", projectName)

	goBunInit(projectUser, projectName)
	goModInit(projectUser, projectName)
	goGitInit(projectUser, projectName, options)
	goOptions(options)

	cmd.Run("code", ".")
}

// goBunInit asigna autor y name al package.json
func goBunInit(projectUser string, projectName string) {
	data, err := os.ReadFile("package.json")
	if err != nil {
		cmd.Danger("No se pudo leer package.json: " + err.Error())
		os.Exit(2)
	}

	var pkg map[string]any
	if err := json.Unmarshal(data, &pkg); err != nil {
		cmd.Danger("Error al parsear package.json: " + err.Error())
		os.Exit(2)
	}

	pkg["name"] = projectName
	pkg["author"] = projectUser

	updated, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		cmd.Danger("Error al serializar package.json: " + err.Error())
		os.Exit(2)
	}

	if err := os.WriteFile("package.json", updated, 0644); err != nil {
		cmd.Danger("No se pudo escribir package.json: " + err.Error())
		os.Exit(2)
	}

	cmd.Runf("Instalando dependencias js", "bun", "install")
	cmd.Success("package.json actualizado con éxito")
}

// goModInit actualiza el nombre del módulo y los imports internos
func goModInit(projectUser string, projectName string) {
	newGoModName := projectUser + "/" + projectName
	oldImport := "donbarrigon/new/internal/"
	newImport := newGoModName + "/internal/"

	cmd.Runf("Actualizando go.mod", "go", "mod", "edit", "-module="+newGoModName)

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			content := string(data)
			if strings.Contains(content, oldImport) {
				newContent := strings.ReplaceAll(content, oldImport, newImport)
				if err := os.WriteFile(path, []byte(newContent), 0644); err != nil {
					return err
				}
				fmt.Println("🔁 Reemplazado import en:", path)
			}
		}
		return nil
	})
	if err != nil {
		cmd.Danger("Error al actualizar imports: " + err.Error())
		os.Exit(2)
	}

	cmd.Runf("Instalando dependencias go", "go", "mod", "tidy")
	cmd.Success("go.mod actualizado.")
}

func goGitInit(projectUser string, projectName string, options []string) {
	cmd.Info("Iniciando la configuracion de git")
	if slices.Contains(options, "fork") {
		cmd.Runf("Renombrando origin a upstream", "git", "remote", "rename", "origin", "upstream")
		cmd.Runf("Agregando cambios al staging", "git", "add", ".")
		cmd.Runf("Realizando commit inicial", "git", "commit", "-m", "feat: initial commit from donbarrigon/new")
	} else {
		cmd.Runf("Eliminando historial de Git existente", "rm", "-rf", ".git")
		cmd.Runf("Inicializando nuevo repositorio Git", "git", "init")
		cmd.Runf("Agregando archivos al staging", "git", "add", ".")
		cmd.Runf("Realizando commit inicial", "git", "commit", "-m", "feat: initial commit from donbarrigon/new")
	}

	cmd.Success("Git de " + projectUser + "/" + projectName + " configurado.")
}

func goOptions(options []string) {
	if !slices.Contains(options, "docs") {
		e := os.RemoveAll("internal/docs")
		if e != nil {
			cmd.Danger("No se pudo eliminar la documentacion: " + e.Error())
			os.Exit(1)
		}
		cmd.Success("Documentacion eliminada")
	}
	if !slices.Contains(options, "tailwind") {
		cmd.Danger("TODO: La opcion tailwind aun no esta implementada.")
	}
}
