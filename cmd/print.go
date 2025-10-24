package cmd

import (
	"donbarrigon/make/cmd/color"
	"fmt"
)

func Success(msg string) {
	fmt.Println(color.Success + "✅ " + msg + color.Reset)
	fmt.Println("")
}

func Info(msg string) {
	fmt.Println(color.Info + "ℹ️ " + msg + color.Reset)
}

func Warning(msg string) {
	fmt.Println(color.Warning + "⚠️ " + msg + color.Reset)
}

func Danger(msg string) {
	fmt.Println("")
	fmt.Println(color.Danger + "💀 " + msg + color.Reset)
	fmt.Println("")
}

func Primary(msg string) {
	fmt.Println("")
	fmt.Println(color.Primary + msg + color.Reset)
	fmt.Println("")
}

func Secondary(msg string) {
	fmt.Println(color.Secondary + msg + color.Reset)
}

func Light(msg string) {
	fmt.Println(color.Light + msg + color.Reset)
}

func Gray(msg string) {
	fmt.Println(color.Gray + msg + color.Reset)
}

func Print(msg string) {
	fmt.Println(msg)
}
