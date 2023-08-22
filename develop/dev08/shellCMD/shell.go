package shellCMD

import (
	"fmt"
	"strings"
)

func Run(inp string) error {
	args := strings.Fields(inp)

	if len(args) < 1 {
		return fmt.Errorf("write a command")
	}
	cmd := args[0]
	switch cmd {
	case "cd":
		Cd(args)
	case "pwd":
		PWD()
	case "echo":
		Echo(args)
	case "kill":
		Kill(args)
	case "ps":
		Ps(args)
	default:
		fmt.Println("unknown command")
		break
	}

	return nil

}
