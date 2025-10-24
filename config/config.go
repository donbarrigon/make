package config

import (
	"donbarrigon/make/cmd"
	"encoding/json"
	"os"
)

func Save(conf map[string]any) {
	data, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		cmd.Danger("Error al serializar config.json: " + err.Error())
		os.Exit(2)
	}

	if err := os.WriteFile("make.config.json", data, 0644); err != nil {
		cmd.Danger("No se pudo escribir config.json: " + err.Error())
		os.Exit(2)
	}
}

func Load() map[string]any {
	data, err := os.ReadFile("make.config.json")
	if err != nil {
		cmd.Danger("No se pudo leer config.json: " + err.Error())
		os.Exit(2)
	}

	var conf map[string]any
	if err := json.Unmarshal(data, &conf); err != nil {
		cmd.Danger("Error al parsear config.json: " + err.Error())
		os.Exit(2)
	}

	return conf
}
