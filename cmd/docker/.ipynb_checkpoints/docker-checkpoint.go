package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func InstallDocker() {
	if !commandExists("docker") {
		fmt.Println("Installing Docker...")
		err := installDocker()
		if err != nil {
			fmt.Println("Error installing Docker:", err)
			return
		}
		fmt.Println("Docker installed successfully!")
	} else {
		fmt.Println("Docker is already installed.")
	}
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func installDocker() error {
	switch runtime.GOOS {
	case "windows":
		return fmt.Errorf("Windows installation not supported")
	case "linux":
		cmd := exec.Command("curl", "-fsSL", "https://get.docker.com", "-o", "get-docker.sh")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}

		cmd = exec.Command("sh", "get-docker.sh")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

func StartDocker() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		fmt.Println("Starting Docker on Windows is not supported")
		return
	case "linux":
		cmd = exec.Command("sudo", "service", "docker", "start")
	default:
		fmt.Println("Unsupported operating system:", runtime.GOOS)
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error starting Docker:", err)
	}
    
}

func StartBashInDocker() {
	cmd := exec.Command("docker", "run", "-it", "ubuntu", "bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error starting bash in Docker container:", err)
	}
}

func StopDocker() {
	cmd := exec.Command("docker", "stop", "my-container")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error stopping Docker container:", err)
	}

	cmd = exec.Command("docker", "rm", "my-container")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error removing Docker container:", err)
	}
}