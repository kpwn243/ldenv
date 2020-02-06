package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Fatalf("You must provide an env file and command to run!")
	}

	commandPath := args[2]
	if len(commandPath) == 0 {
		log.Fatalf("You must specify a command to run")
	}
	commandParams := args[3:]

	filePath := args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open the specified file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if len(line) != 0 && !strings.HasPrefix(line, "#") && strings.Contains(line, "=") {
			parts := strings.Split(scanner.Text(), "=")
			err = os.Setenv(parts[0], parts[1])
			if err != nil {
				log.Fatalf("Failed to set environment variable with key: %s", parts[0])
			}
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("A failure occurred while reading the file: %v", err)
	}

	cmd := exec.Command(commandPath, commandParams...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run the command specified: %v", err)
	}
}
