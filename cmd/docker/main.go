package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	InstallDocker()

	fmt.Println("Starting Docker...")
	StartDocker()

	startZsh()
}

func startZsh() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("zsh> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Exiting Zsh shell...")
			break
		}

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		err = cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error executing command:", err)
		}
	}
}
