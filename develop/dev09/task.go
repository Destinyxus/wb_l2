package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"dev09/pkg"
)

func main() {

	fmt.Print("Enter a command: ")
	reader := bufio.NewReader(os.Stdin)
	command, _ := reader.ReadString('\n')

	s := strings.Fields(command)
	if len(s) < 2 {
		fmt.Println("wget: missing URL")
		return
	}

	if s[0] == "wget" && (strings.HasPrefix(s[1], "http") || strings.HasPrefix(s[1], "https")) {
		fmt.Println("hi")
		wget := pkg.NewWget()
		if err := wget.GetSite(s[1]); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("write a link properly")
	}
}
