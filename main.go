package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	destroy := false

	if len(os.Args) > 1 && os.Args[1] == "--destroy=true" {
		destroy = true
	} else if len(os.Args) > 1 && os.Args[1] == "--destroy=false" {
		destroy = false
	}

	if err := os.Chdir("terraform"); err != nil {
		fmt.Printf("Failed to change directory: %v\n", err)
		return
	}

	fmt.Println("Initializing Terraform...")
	if err := runTerraform("init"); err != nil {
		fmt.Println("Error initializing Terraform:", err)
		return
	}

	fmt.Println("Applying Terraform...")
	if err := runTerraform("apply", "-auto-approve"); err != nil {
		fmt.Println("Error applying Terraform:", err)
		return
	}

	fmt.Println("Terraform applied successfully.")

	if destroy {
		destroyTerraform()
	} else {
		fmt.Println("Terraform resources will not be destroyed.")
	}
}

func runTerraform(args ...string) error {
	cmd := exec.Command("terraform", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func destroyTerraform() {
	fmt.Println("Destroying Terraform resources...")
	if err := runTerraform("destroy", "-auto-approve"); err != nil {
		fmt.Println("Error destroying Terraform resources:", err)
		return
	}
	fmt.Println("Terraform destroyed successfully.")
}
