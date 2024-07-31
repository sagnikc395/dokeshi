package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	serveFlag := flag.Bool("s", false, "serve a local server to view files")

	flag.Parse()

	argNums := len(os.Args)

	if argNums < 2 {
		cli.Run()
	} else {
		if *serveFlag {
			cli.Serve()
		} else {
			fmt.Println("Please check the arguments.")
		}
	}
}
