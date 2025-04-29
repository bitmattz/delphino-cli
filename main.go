package main

import (
	"fmt"
)

func main() {

	// CREATE MENU
	fmt.Println(`
$$$$$$$\  $$$$$$$$\ $$\       $$$$$$$\  $$\   $$\ $$$$$$\ $$\   $$\  $$$$$$\  
$$  __$$\ $$  _____|$$ |      $$  __$$\ $$ |  $$ |\_$$  _|$$$\  $$ |$$  __$$\ 
$$ |  $$ |$$ |      $$ |      $$ |  $$ |$$ |  $$ |  $$ |  $$$$\ $$ |$$ /  $$ |
$$ |  $$ |$$$$$\    $$ |      $$$$$$$  |$$$$$$$$ |  $$ |  $$ $$\$$ |$$ |  $$ |
$$ |  $$ |$$  __|   $$ |      $$  ____/ $$  __$$ |  $$ |  $$ \$$$$ |$$ |  $$ |
$$ |  $$ |$$ |      $$ |      $$ |      $$ |  $$ |  $$ |  $$ |\$$$ |$$ |  $$ |
$$$$$$$  |$$$$$$$$\ $$$$$$$$\ $$ |      $$ |  $$ |$$$$$$\ $$ | \$$ | $$$$$$  |
\_______/ \________|\________|\__|      \__|  \__|\______|\__|  \__| \______/ 
																			  
`)
	fmt.Println("Welcome to the Package Manager!")
	fmt.Println("This program allows you to install preset packages and create folders on your system.")
	fmt.Println("\n")
	fmt.Println("Please choose an option:")
	fmt.Println("1. Install a package")
	fmt.Println("2. Create a folder")
	fmt.Println("3. Uninstall preset packages")
	fmt.Println("4. Delete folders")
	fmt.Println("5. Exit")

	// READ JSON

	// SEGMENTATION

	// INSTALATION

	//UNINSTALLATION

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter the name of the package to install: ")
	// packageName, _ := reader.ReadString('\n')
	// packageName = strings.TrimSpace(packageName)

	// if packageName == "" {
	// 	fmt.Println("Package name cannot be empty.")
	// 	return
	// }

	// cmd := exec.Command("sudo", "apt-get", "install", "-y", packageName)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// fmt.Println("Installing package:", packageName)
	// if err := cmd.Run(); err != nil {
	// 	fmt.Println("Error installing package:", err)
	// } else {
	// 	fmt.Println("Package installed successfully.")
	// }
}
