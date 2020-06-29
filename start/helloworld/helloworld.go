package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "this is null")
}

func main() {
	flag.Parse()
	hello(name)
}

func hello(name string) {
	fmt.Printf("hello,%s! \n", name)
}
