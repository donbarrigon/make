package cmd

import (
	"donbarrigon/make/pkg/str"
	"os"
)

type Name struct {
	Snake   string
	Camel   string
	Pascal  string
	Kebab   string
	Var     string
	Snakep  string
	Camelp  string
	Pascalp string
	Kebabp  string
}

// comandos que no se ejecutan en grupo
var CommandList = []string{
	"help",
	"project",
	"merge:upstream",
	"run",
	"build",
}

// comandos que se ejecutan en grupo
var TemplateList = []string{
	"model",
	"migration",
	"seed",
	"repository",
	"resource",
	"view",
	"page",
	"component",
	"controller",
	"middleware",
	"policy",
	"route",
	"service",
	"validator",
	"db",
	"ui",
	"handler",
	"api",
	"mvc",
}

func GetCommands() []string {

	if len(os.Args) < 2 {
		return []string{"help"}
	}

	command := os.Args[1]
	for _, c := range CommandList {
		if c == command {
			return []string{command}
		}
	}

	args := os.Args[1 : len(os.Args)-1]
	if len(args) == 0 {
		Danger("El comando " + command + " nesesita almenos un argumento")
		return []string{"help"}
	}
	for _, a := range args {
		exist := false
		for _, t := range TemplateList {
			if t == a {
				exist = true
				break
			}
		}
		if !exist {
			return []string{"[" + a + "]"}
		}
	}

	return args
}

func GetName() Name {
	n := os.Args[len(os.Args)-1]
	return Name{
		Snake:   str.ToSnakeCase(n),
		Camel:   str.ToCamelCase(n),
		Pascal:  str.ToPascalCase(n),
		Kebab:   str.ToKebabCase(n),
		Var:     str.ToVar(n),
		Snakep:  str.PluralizeIdentifier(str.ToSnakeCase(n)),
		Camelp:  str.PluralizeIdentifier(str.ToSnakeCase(n)),
		Pascalp: str.PluralizeIdentifier(str.ToSnakeCase(n)),
		Kebabp:  str.PluralizeIdentifier(str.ToSnakeCase(n)),
	}
}
