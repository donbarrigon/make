package cmd

import (
	"bufio"
	"donbarrigon/make/cmd/color"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Prompt(msg string, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg + " [" + color.Secondary + defaultValue + color.Reset + "]: ")
	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(value)

	if value == "" {
		return defaultValue
	}
	return value
}

func PromptSelect(msg string, defaultValue int, options ...string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [%s%s%s]:\n", msg, color.Secondary, options[defaultValue], color.Reset)
	for i, option := range options {
		fmt.Printf("[%d] => %s\n", i+1, option)
	}

	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.ToLower(strings.TrimSpace(choiceStr))

	choice := defaultValue
	if choiceStr != "" {
		c, e := strconv.Atoi(strings.TrimSpace(choiceStr))
		if e == nil {
			if c < 1 || c > len(options) {
				Danger("!FIJESE, el valor [" + choiceStr + " ] esta fuera del rango")
				os.Exit(2)
			}
			choice = c - 1
		} else {
			c := -1
			for i, option := range options {
				if strings.ToLower(option) == choiceStr {
					c = i
					break
				}
			}
			choice = c
		}
	}

	return options[choice]
}

func PromptConfirm(msg string, defaultValue bool) bool {
	reader := bufio.NewReader(os.Stdin)

	var hint string
	if defaultValue {
		hint = "[" + color.Secondary + "Y" + color.Reset + "/n]"
	} else {
		hint = "[y/" + color.Secondary + "N" + color.Reset + "]"
	}

	fmt.Printf("%s %s %s: ", msg, color.Secondary, hint+color.Reset)

	value, _ := reader.ReadString('\n')
	value = strings.ToLower(strings.TrimSpace(value))

	if value == "" {
		return defaultValue
	}
	switch value {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	default:
		Danger("No veo cual es la dificultad que tenes para poner Y o N, esto que pones [" + value + "] esta mal.")
		os.Exit(2)
		return false
	}
}
