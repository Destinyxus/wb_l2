package shellCMD

import (
	"fmt"
	"log"
	"os"
)

func PWD() {
	currPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("current path is %s\n", currPath)
}
