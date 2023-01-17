package logger

import (
	"fmt"
	"os"
)

func println(msg string, colorCode string) {
	fmt.Fprintf(os.Stdout, "\033["+colorCode+"m%s\033[0m\n", msg)
}

func Debug(msg string) {
	println(msg, "34")
}

func Info(msg string) {
	println(msg, "32")
}

func Warn(msg string) {
	println(msg, "33")
}

func Error(msg string) {
	println(msg, "31")
}
