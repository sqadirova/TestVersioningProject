package main

import (
	"flag"
	"fmt"
	"log"
)

//var gitCommit string
var Version = "development"

func printVersion() {
	log.Printf("Current build version: %s", Version)
}

func main() {
	versionFlag := flag.Bool("v", false, "Print the current version and exit")
	flag.Parse()

	switch {
	case *versionFlag:
		printVersion()
		return
	}

	//fmt.Println("Version:\t", Version)

	fmt.Println("Hello, world!!!")
}
