package main

import (
	"os"
	"push/algo"
)

func main() {
	args := os.Args[1:]

	algo.Push_swap(args)
}
