# Docker

Welcome to the Docker project within Godoc. This project focuses on Docker management and interaction using GoLang.

## Overview

The Docker project provides functionalities for installing Docker, starting Docker containers, and interacting with Dockerized environments using GoLang.

## Running the Project

To run the Docker project, follow these steps:

1. Ensure that GoLang is installed on your system. You can download and install GoLang from the [official website](https://golang.org/dl/).

2. Navigate to the `docker` directory in your terminal.

3. Run the main.go file along with the docker.go file using the following command:

```
go run main.go docker.go
```

This command will execute the main Go program, which pulls the Kali Linux Docker image and starts a Docker container. Additionally, it provides a Zsh shell interface for interacting with the containerized environment.

## Requirements for Windows

If you are running the Docker project on Windows, ensure the following:

- Docker Desktop is installed on your system. You can download Docker Desktop from the [official website](https://www.docker.com/products/docker-desktop).
- Docker Desktop is running and accessible from the command line.
- GoLang is installed on your system, and the environment variables are configured correctly.

Please note that Docker Desktop provides a Docker CLI that can be used from any shell, including PowerShell and Command Prompt, to interact with Docker containers.

