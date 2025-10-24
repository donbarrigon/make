package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Run(name string, args ...string) {
	msg := fmt.Sprintf("Ejecutando: %s %s", name, strings.Join(args, " "))
	Info(msg)
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func Runf(msg string, name string, args ...string) {
	Info(msg)
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
