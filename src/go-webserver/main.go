package main

import (
	"flag"
	"fmt"
)

func main() {
	ver := flag.Bool("version", false, "show version info")
	flag.Parse()
	if *ver {
		fmt.Println(verinfo())
		return
	}
}
