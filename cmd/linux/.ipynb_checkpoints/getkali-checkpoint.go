package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Pulling Kali Linux Docker image...")
	cmd := exec.Command("docker", "pull", "kalilinux/kali-linux-docker")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error pulling Kali Linux Docker image:", err)
	}
}
