package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/aushafy/cryptonice-api/app"
)

func main() {
	// Check condition of cryptonice scanner already on the server or not
	// find cryptonice binary on the right path

	fmt.Printf("################################################\n")
	fmt.Printf("### Checking requirements before running apps ###\n")
	fmt.Printf("################################################\n")

	cryptonicePath, err := exec.LookPath("cryptonice")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	} else {
		fmt.Printf("Checking completed\nFound 1 requirement of cryptonice binary files on %s \n\n", cryptonicePath)
	}

	app.StartApplication()
}
