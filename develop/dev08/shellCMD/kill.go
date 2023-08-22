package shellCMD

import (
	"fmt"
	"log"
	"strconv"
	"syscall"
)

func Kill(args []string) {

	if len(args) == 1 {
		fmt.Println("kill usage: kill 'pid'")
		return
	}
	k, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}
	err = syscall.Kill(k, syscall.SIGTERM)
	if err != nil {
		log.Fatal(err)
	}

}
