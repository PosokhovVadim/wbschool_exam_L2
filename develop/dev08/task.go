package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
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

func cd(args string) {
	if len(args) < 1 {
		fmt.Println("cd: missing argument")
		return
	}
	err := os.Chdir(args)
	if err != nil {
		fmt.Println("cd:", err)
	}

}

func pwd() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("pwd:", err)
	}
	fmt.Println(path)
}

func echo(args []string) {
	if len(args) >= 1 {

		fmt.Println(strings.Join(args, " "))
	}
}

func kill(arg string) {
	pid, err := strconv.Atoi(arg)

	if err != nil {
		fmt.Println("kill:", err)
	}
	err = syscall.Kill(pid, syscall.SIGKILL)
	if err != nil {
		fmt.Println("kill:", err)
	}
}

func ps() {
	cmd := exec.Command("ps", "aux")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("ps:", err)
	}
}
func runPipeline(commands [][]string) error {
	var cmds []*exec.Cmd

	for _, args := range commands {
		cmd := exec.Command(args[0], args[1:]...)
		cmds = append(cmds, cmd)
	}

	for i := 0; i < len(cmds)-1; i++ {
		stdoutPipe, err := cmds[i].StdoutPipe()
		if err != nil {
			return err
		}
		cmds[i+1].Stdin = stdoutPipe
	}

	cmds[len(cmds)-1].Stdout = os.Stdout

	for _, cmd := range cmds {
		if err := cmd.Start(); err != nil {
			return err
		}
	}

	for _, cmd := range cmds {
		if err := cmd.Wait(); err != nil {
			return err
		}
	}
	return nil
}

func run() error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	var input string
	for scanner.Scan() {
		input = scanner.Text()

		if strings.Contains(input, "|") {
			commands := strings.Split(input, "|")
			var cmdList [][]string

			for _, v := range commands {
				cmdList = append(cmdList, strings.Fields(strings.TrimSpace(v)))
			}

			err := runPipeline(cmdList)
			if err != nil {
				return err
			}

		} else {

			args := strings.Fields(strings.TrimSpace(input))

			if len(args) == 0 {
				continue
			}

			switch args[0] {
			case "cd":
				cd(args[1])
			case "pwd":
				pwd()
			case "echo":
				echo(args[1:])
			case "kill":
				kill(args[1])
			case "ps":
				ps()
			case "\\quit":
				return nil
			default:
				cmd := exec.Command(args[0], args[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					fmt.Printf("error: %v\n", err)
				}
			}
		}
		fmt.Print("> ")
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

}
