package utils

import (
	"log"

	"github.com/fatih/color"
)

func Info(msg string) {
	infoColor := color.New(color.FgGreen).SprintFunc()
	log.Println(infoColor("[OK] " + msg))
}

func Warn(msg string) {
	warnColor := color.New(color.FgHiYellow).SprintFunc()
	log.Println(warnColor("[WARNING] " + msg))
}

func Error(msg string) {
	errorColor := color.New(color.FgRed).SprintFunc()
	log.Println(errorColor("[ERROR] " + msg))
}
func ErrorFatal(msg string) {
	errorColor := color.New(color.FgRed).SprintFunc()
	log.Fatalln(errorColor("[FATAL]" + msg))
}

func InfoBlue(msg string) {
	errorColor := color.New(color.FgBlue).SprintFunc()
	log.Println(errorColor(msg))
}
