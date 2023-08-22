package shellCMD

import (
	"fmt"
	"os"
	"os/exec"
)

func Ps(args []string) {
	if len(args) == 1 {
		fmt.Println("ps usage: ps aux`")
		return
	}

	cmd := exec.Command("ps", "aux")
	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Print(string(output))
}
