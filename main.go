package main

import (
	"fmt"
)

var Version = "development"

//func printVersion() {
//	log.Printf("Current build version: %s", Version)
//}

func main() {
	//versionFlag := flag.Bool("v", false, "Print the current version and exit")
	//flag.Parse()
	//
	//switch {
	//case *versionFlag:
	//	printVersion()
	//	return
	//}
	// continue with other stuff
	fmt.Println("Hello, world!!!")
	//fmt.Println(versionFlag)
	fmt.Println("Version:\t", Version)
}

//var gitCommit string
//
//func printVersion() {
//	log.Printf("Current build version: %s", gitCommit)
//}
//
//func main() {
//	versionFlag := flag.Bool("v", false, "Print the current version and exit")
//	flag.Parse()
//
//	switch {
//	case *versionFlag:
//		printVersion()
//		return
//	}
//	// continue with other stuff
//	fmt.Println("Hello, world!!!")
//	fmt.Println(versionFlag)
//	fmt.Println(gitCommit)
//}
