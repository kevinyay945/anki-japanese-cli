package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Running Anki Connect tests...")
	fmt.Println("Make sure Anki is running with AnkiConnect plugin installed")
	fmt.Println("and a deck named '日文學習' exists.")
	fmt.Println()

	cmd := exec.Command("go", "test", "-v", "./internal/anki")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Tests failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("All tests passed!")
}
