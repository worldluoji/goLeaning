package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Hello, world!!!")
	for _, param := range args {
		fmt.Println(param)
	}
	os.Exit(0)
}
