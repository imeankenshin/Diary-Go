package ui

import "fmt"

// The funcs bellow output a colorized string
func Gray (format string) {
	fmt.Print("\033[30m"+format+"\033[0m")
}

func Red (format string) {
	fmt.Print("\033[31m"+format+"\033[0m")
}

func Green (format string) {
	fmt.Print("\033[32m"+format+"\033[0m")
}

func Yellow (format string) {
	fmt.Print("\033[33m"+format+"\033[0m")
}

func Blue (format string) {
	fmt.Print("\033[34m"+format+"\033[0m")
}

func Pink (format string) {
	fmt.Print("\033[35m"+format+"\033[0m")
}

func Sky (format string) {
	fmt.Print("\033[36m"+format+"\033[0m")
}

// The funcs bellow output a colorized string
func Grayln (format string) {
	fmt.Println("\033[30m"+format+"\033[0m")
}

func Redln (format string) {
	fmt.Println("\033[31m"+format+"\033[0m")
}

func Greenln (format string) {
	fmt.Println("\033[32m"+format+"\033[0m")
}

func Yellowln (format string) {
	fmt.Println("\033[33m"+format+"\033[0m")
}

func Blueln (format string) {
	fmt.Println("\033[34m"+format+"\033[0m")
}

func Pinkln (format string) {
	fmt.Println("\033[35m"+format+"\033[0m")
}

func Skyln (format string) {
	fmt.Println("\033[36m"+format+"\033[0m")
}
