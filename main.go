package main

import (
	"donbarrigon/make/cmd"
	"donbarrigon/make/project"
	"fmt"
	"os"
)

func main() {

	commnds := cmd.GetCommands()
	fmt.Println(commnds)
	for _, c := range commnds {
		switch c {
		case "help":
			cmd.Help()
		case "project":
			project.New()
		case "merge:upstream":
			cmd.Info("merge:upstream")
		case "run":
			cmd.Info("run")
		case "build":
			cmd.Info("build")
		case "model":
			cmd.Info("model")
		case "migration":
			cmd.Info("migration")
		case "seed":
			cmd.Info("seed")
		case "repository":
			cmd.Info("repository")
		case "resource":
			cmd.Info("resource")
		case "view":
			cmd.Info("view")
		case "page":
			cmd.Info("page")
		case "component":
			cmd.Info("component")
		case "controller":
			cmd.Info("controller")
		case "middleware":
			cmd.Info("middleware")
		case "policy":
			cmd.Info("policy")
		case "route":
			cmd.Info("route")
		case "validator":
			cmd.Info("validator")
		case "db":
			cmd.Info("db")
		case "ui":
			cmd.Info("ui")
		case "handler":
			cmd.Info("handler")
		case "api":
			cmd.Info("api")
		case "mvc":
			cmd.Info("mvc")
		default:
			cmd.Help()
			cmd.Danger("El comando " + c + " no existe")
			fmt.Println("")
			os.Exit(127)
		}
	}
	fmt.Println("")
}
