package command

import (
	"fmt"
	"log"
	"os/exec"
)

// RunCommand exec command line
func RunCommand(arg1, arg2 int) string {

	command := fmt.Sprintf("go run ./src %d %d", arg1, arg2)

	cmd := exec.Command(command)

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	b, _ := cmd.Output()
	return string(b)
}
