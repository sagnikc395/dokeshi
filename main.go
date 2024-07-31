package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sagnikc395/dokeshi/dokeshi"
)

func main() {
	serveFlag := flag.Bool("s", false, "serve a local server to view files")

	flag.Parse()

	argNums := len(os.Args)

	if argNums < 2 {
		dokeshi.Run()
	} else {
		if *serveFlag {
			dokeshi.Serve()
		} else {
			fmt.Println("Please check the arguments.")
		}
	}
}
