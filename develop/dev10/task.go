package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"dev10/telnet"
)

var (
	flagTimeout string
)

func main() {
	flag.StringVar(&flagTimeout, "timeout", "10s", "Connection timeout duration")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("telnet usage:...")
		return
	}

	addr, ok := telnet.Concat(flag.Arg(0), flag.Arg(1))
	if !ok {
		fmt.Println("provide a correct address")
		return
	}
	timeout, err := time.ParseDuration(flagTimeout)
	client := telnet.NewTelnet(addr, time.Duration(timeout.Seconds()))
	err = client.Run()
	if err != nil {
		log.Fatal(err)

	}

}
