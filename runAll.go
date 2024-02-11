package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	scripts := []string{"retrieveWorking.go", "toSingleLine.go", "processEscaped.go"}

	for _, script := range scripts {
		// Compile the Go script
		buildCmd := exec.Command("go", "build", script)
		err := buildCmd.Run()
		if err != nil {
			log.Fatalf("Failed to compile %s: %v", script, err)
		}

		// Assuming the output binary has the same name as the script but without .go extension
		binary := script[:len(script)-3]

		// Execute the compiled binary
		fmt.Printf("Running %s...\n", binary)
		runCmd := exec.Command("./" + binary)
		runCmd.Stdout = os.Stdout // Forward stdout to the main process' stdout
		runCmd.Stderr = os.Stderr // Forward stderr to the main process' stderr
		err = runCmd.Run()
		if err != nil {
			log.Fatalf("Failed to run %s: %v", binary, err)
		}

		fmt.Printf("%s completed\n", binary)

		// Optionally, remove the binary after execution to clean up, if desired
		err = exec.Command("rm", binary).Run()
		if err != nil {
			log.Printf("Failed to remove binary %s: %v", binary, err)
		}
	}
}
