package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		commandHandler(scanner.Text())
	}

}

func commandHandler(stringCommand string) {

	switch strings.Split(stringCommand, " ")[0] {

	case "cd":
		chDirCommand(stringCommand)

	case "pwd":
		pwdCommand()

	case "echo":
		echoCommand(stringCommand)

	case "kill":
		killPsCommand(stringCommand)

	case "ps":
		psCommand()

	case `\quit`:
		exitCommand()

	default:
		fmt.Println("Invalid command")
	}
}

func chDirCommand(stringCommand string) {
	err := os.Chdir(strings.Replace(stringCommand, "cd ", "", 1))
	if err != nil {
		fmt.Println(err)
	}
}

func pwdCommand() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
}

func echoCommand(stringCommand string) {
	fmt.Println(strings.Replace(stringCommand, "echo ", "", 1))
}

func killPsCommand(stringCommand string) {
	pid, err := strconv.Atoi(strings.Replace(stringCommand, "kill ", "", 1))
	if err != nil {
		fmt.Println(err)
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err)
	}
	proc.Kill()
}

func psCommand() {
	sliceProc, _ := ps.Processes()
	for _, proc := range sliceProc {
		fmt.Printf("Process name: %v process id: %v\n", proc.Executable(), proc.Pid())
	}
}

func exitCommand() {
	fmt.Println("Exit")
	os.Exit(0)
}
