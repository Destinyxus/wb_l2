package shellCMD

import (
	"fmt"
	"log"
	"os"
)

func Cd(dir []string) {
	if len(dir) == 1 {
		hd, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		if err := os.Chdir(hd); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("moved into: %s directory\n", hd)

		return
	} else if len(dir) > 1 {
		if err := os.Chdir(dir[1]); err != nil {
			return
		}
		fmt.Printf("moved into: %s directory\n", dir[1])

	}

}
