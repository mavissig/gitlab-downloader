package formating

import (
	"fmt"
	"strings"
)

/*
-------------------------------------------------------
		        FORMATING COLORS
-------------------------------------------------------
*/

const (
	nc = "\033[0m"

	red    = "\033[0;31m"
	green  = "\033[0;32m"
	yellow = "\033[0;33m"

	redBold    = "\033[1;31m"
	greenBold  = "\033[1;32m"
	yellowBold = "\033[1;33m"

	redBG    = "\033[41m"
	greenBG  = "\033[42m"
	yellowBG = "\033[43m"

	redBG_black    = "\033[41;30m"
	greenBG_black  = "\033[42;30m"
	yellowBG_black = "\033[43;30m"
)

/*
-------------------------------------------------------
		        FORMATING SUFFIX FUNCTIONS
-------------------------------------------------------
*/

func suffixInfo() string {
	return fmt.Sprintf("%sINFO%s", yellow, nc)
}
func suffixError() string {
	return fmt.Sprintf("%sERROR%s", red, nc)
}

func suffixSuccess() string {
	return fmt.Sprintf("%sSUCCESS%s", green, nc)
}

func suffixWarning() string {
	return fmt.Sprintf("%sWARNING%s", yellow, nc)
}

/*
-------------------------------------------------------
		        FORMATING TEXT FUNCTIONS
-------------------------------------------------------
*/

func YellowText(text string) string {
	return fmt.Sprintf("%s%s%s", yellow, text, nc)
}

func RedText(text string) string {
	return fmt.Sprintf("%s%s%s", red, text, nc)
}

func GreenText(text string) string {
	return fmt.Sprintf("%s%s%s", green, text, nc)
}

func YellowBoldText(text string) string {
	return fmt.Sprintf("%s%s%s", yellowBold, text, nc)
}

func RedBoldText(text string) string {
	return fmt.Sprintf("%s%s%s", redBold, text, nc)
}

func GreenBoldText(text string) string {
	return fmt.Sprintf("%s%s%s", greenBold, text, nc)
}

/*
-------------------------------------------------------
		        FORMATING LOG FUNCTIONS
-------------------------------------------------------
*/

func LogInfo(msg ...string) string {
	return strings.Join(append([]string{suffixInfo()}, msg...), " ")
}

func LogError(msg ...string) string {
	return strings.Join(append([]string{suffixError()}, msg...), " ")
}

func LogSuccess(msg string) string {
	return fmt.Sprintf("%sSUCCESS%s %s", greenBold, nc, msg)
}

func LogWarning(msg string) string {
	return fmt.Sprintf("%sWARNING%s %s", yellowBold, nc, msg)
}
