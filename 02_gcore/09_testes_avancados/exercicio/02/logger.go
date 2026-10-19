package main

import "fmt"

func FormatadorComPlus(servico, msg string) string {

	return "[" + servico + "] " + msg
}

func FormatadorComSprintf(servico, msg string) string {
	return fmt.Sprintf("[%v] %v", servico, msg)
}