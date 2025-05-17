package main

import (
	"os"
	"strconv"
)

func main() {
	cmd := os.Args[1]
	a, _ := strconv.Atoi(os.Args[2])
	b, _ := strconv.Atoi(os.Args[3])
}
