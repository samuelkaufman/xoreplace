package main

import (
	"log"
	"os"

	"github.com/samuelkaufman/xoreplace/pkg/xoreplace"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("USAGE: xoreplace PATTERN")
	}
	replacer := xoreplace.New(os.Args[1], nil)
	replacer.Run()
}
