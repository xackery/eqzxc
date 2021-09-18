package main

import (
	"fmt"
	"os"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("failed:", err)
		os.Exit(1)
	}
}

func run() error {
	return nil
}
