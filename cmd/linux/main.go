package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Starting Kali Linux Docker container...")
	startKali()

	fmt.Println("Starting Zsh shell...")
	startZsh()
}

func startKali() {
	cmd := exec.Command("docker", "run", "-it", "--name", "kali-container", "kalilinux/kali-linux-docker")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error starting Kali Linux Docker container:", err)
	}
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
