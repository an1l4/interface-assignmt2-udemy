package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("failed to open the file")
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
