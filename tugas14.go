package main

import (
	"fmt"
	"os/exec"
)

func main() {

	var hasil1, _ = exec.Command("git", "config", "user.name").Output()
	fmt.Printf(" -> git config user.name\n%s\n", string(hasil1))

	var hasil2, _ = exec.Command("git", "config", "user.email").Output()
	fmt.Printf(" -> git config user.email\n%s\n", string(hasil2))

	var hasil3, _ = exec.Command("git", "--version").Output()
	fmt.Printf(" -> git --version\n%s\n", string(hasil3))

}
