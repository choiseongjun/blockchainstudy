package main

import (
	"fmt"
	"golangblock/cli"
	"os"
)

func usage() {
	fmt.Printf("Welcome to 노마드 코인\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("explorer:	Start the HTML Explorer\n")
	fmt.Printf("rest:		Start the REST API (recommended)\n\n")
	os.Exit(0)
}
func main() {
	//go explorer.Start(3000)
	//rest.Start(4000)
	//fmt.Println(os.Args)
	cli.Start()
}
