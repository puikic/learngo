package main

import (
	"fmt"
	"os"
)

func main() {
	// go run hello.go cpq
	fmt.Println(os.Args)
	// 输出：cpq
	fmt.Println("hello world")
	os.Exit(0)
}
