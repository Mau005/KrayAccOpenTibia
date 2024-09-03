package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
)

var ListLogInfo []string

func Info(msg ...string) {

	infoColor := color.New(color.FgGreen).SprintFunc()
	log.Println(infoColor(uniteText("[OK]", msg)))
}

func Warn(msg ...string) {
	warnColor := color.New(color.FgHiYellow).SprintFunc()
	log.Println(warnColor(uniteText("[WARNING]", msg)))
}

func WarnLog(msg ...string) {
	warnColor := color.New(color.FgHiYellow).SprintFunc()
	log.Println(warnColor(uniteText("[ALERT SECURITY]", msg)))
}

func Error(msg ...string) {
	errorColor := color.New(color.FgRed).SprintFunc()
	log.Println(errorColor(uniteText("[ERROR]", msg)))
}
func ErrorFatal(msg ...string) {
	errorColor := color.New(color.FgRed).SprintFunc()
	log.Fatalln(errorColor(uniteText("[FATAL]", msg)))
}

func InfoBlue(msg ...string) {
	errorColor := color.New(color.FgBlue).SprintFunc()
	log.Println(errorColor(msg))
}

func InfoBlueNotLog(msg ...string) {
	errorColor := color.New(color.FgBlue).SprintFunc()
	fmt.Println(errorColor(msg))
}

func uniteText(target string, msg []string) string {
	msgComplex := []string{target}

	// Agregar los elementos de msg a msgComplex
	msgComplex = append(msgComplex, msg...)

	// Unir todos los elementos en una sola cadena con espacios entre ellos
	return strings.Join(msgComplex, " ")
}
