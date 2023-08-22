package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"dev08/shellCMD"
)

func main() {
	for {
		fmt.Print("Enter a command: ")
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "\\quit" {
			fmt.Println("Exiting the shell.")
			break
		}

		err := shellCMD.Run(command)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
