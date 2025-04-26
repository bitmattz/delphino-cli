package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name of the package to install: ")
	packageName, _ := reader.ReadString('\n')
	packageName = strings.TrimSpace(packageName)

	if packageName == "" {
		fmt.Println("Package name cannot be empty.")
		return
	}

	cmd := exec.Command("sudo", "apt-get", "install", "-y", packageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Installing package:", packageName)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error installing package:", err)
	} else {
		fmt.Println("Package installed successfully.")
	}
}
