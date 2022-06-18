package main

import (
	"golangblock/explorer"
	"golangblock/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
