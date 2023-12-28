package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	myShell()
}

func myShell() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("[myShell]: ")
		if scanner.Scan() {
			line := scanner.Text()
			str := strings.Split(line, " ")
			command := str[0]
			args := str[1:]
			switch command {
			case "pwd":
				pwdCommand(args)
			case "echo":
				fmt.Println(strings.Join(args, " "))
			case "cd":
				cdCommand(args)
			case "ps":
				psCommand(args)
			case "kill":
				killCommand(args)
			case "exec":
				execCommand(args)
			case "q", "quit":
				os.Exit(0)
			default:
				fmt.Fprintln(os.Stderr, "unknown command:", command)
			}
		}
	}
}

func pwdCommand(args []string) {
	if len(args) != 0 {
		fmt.Fprintln(os.Stderr, "pwd: too many arguments")
		return
	}
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(dir)
	}
}

func cdCommand(args []string) {
	if len(args) == 0 {
		homedir, _ := os.UserHomeDir()
		if err := os.Chdir(homedir); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	} else if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "cd: too many arguments")
	} else {
		if err := os.Chdir(args[0]); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func psCommand(args []string) {
	if len(args) != 0 {
		fmt.Fprintln(os.Stderr, "ps: too many arguments")
		return
	}
	processList, err := ps.Processes() // библиотека go-ps, код взят из stackoverflow.com
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("PID\tEXECUTABLE NAME")
	for x := range processList {
		var process ps.Process
		process = processList[x]
		fmt.Printf("%d\t%s\n", process.Pid(), process.Executable())
	}
}

func killCommand(args []string) {
	var pids []int
	for i := 0; i < len(args); i++ {
		pidnum, err := strconv.Atoi(args[i])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		pids = append(pids, pidnum)
	}
	for _, pid := range pids {
		procces, err := os.FindProcess(pid)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if err := procces.Kill(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execCommand(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "error: missing command")
		return
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
