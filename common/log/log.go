package log

import (
	"fmt"
	"time"
)

const colorNone = "\033[0m"
const colorRed = "\033[0;31m"
const colorYellow = "\033[0;33m"
const colorBlue = "\033[0;34m"
const colorMagenta = "\033[0;35m"

func Formatter(header string, content string, color string, colorAll bool) string {
	current_date := time.Now()
	if colorAll {
		return fmt.Sprintf("%s%02d:%02d:%02d - [%s]:%s %s", color, current_date.Hour(), current_date.Minute(), current_date.Second(), header, content, colorNone)
	} else {
		return fmt.Sprintf("%s%02d:%02d:%02d - [%s]:%s %s", color, current_date.Hour(), current_date.Minute(), current_date.Second(), header, colorNone, content)
	}
}

func Warning(message string, a ...any) {
	println(Formatter("WARNING", fmt.Sprintf(message, a...), colorYellow, true))
}

func Info(message string, a ...any) {
	println(Formatter("INFO", fmt.Sprintf(message, a...), colorBlue, false))
}

func Error(message string, a ...any) {
	println(Formatter("ERROR", fmt.Sprintf(message, a...), colorRed, true))
}

func Fatal(message string, a ...any) {
	formatted := fmt.Sprintf(message, a...)
	println(Formatter("FATAL", formatted, colorRed, true))
	panic(formatted)
}

func Debug(message string, a ...any) {
	println(Formatter("DEBUG", fmt.Sprintf(message, a...), colorMagenta, true))
}
