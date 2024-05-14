package shell

import "fmt"

const (
	Reset = "\033[0m"
	red = "\033[31m"
	cyan = "\033[36m"
	bold = "\033[1m"
)

func formatString(ansiFormat, inputStr string) string {
	return fmt.Sprintf("%s%s%s", ansiFormat, inputStr, Reset)
}

func Red(inputStr string) string {
	return formatString(red, inputStr)
}

func Cyan(inputStr string) string {
	return formatString(cyan, inputStr)
}

func Bold(inputStr string) string {
	return formatString(bold, inputStr)
}
