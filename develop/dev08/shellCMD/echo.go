package shellCMD

import (
	"fmt"
	"strings"
)

func Echo(args []string) {

	if len(args) == 1 {
		fmt.Println("")
		return
	}
	finalOutput := make([]string, 0, len(args[1:]))

	for i := 1; i < len(args); i++ {
		finalOutput = append(finalOutput, args[i])
	}

	fmt.Println(strings.Join(finalOutput, " "))
}
