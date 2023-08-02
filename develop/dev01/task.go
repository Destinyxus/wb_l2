package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func main() {
	getTime()
}

func getTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		fmt.Fprintf(os.Stderr, "error occured: %v", err)
		os.Exit(1)
	}

	fmt.Println(time)

}
