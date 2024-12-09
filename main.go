package main

import (
	"fmt"
	"os"
)

func main() {
	// [errcheck]: Error return value of `os.Open` is not checked (errcheck)
	file, _ := os.Open("hogetest.txt")

	unusedSlice := make([]int, 10)
	_, _ = fmt.Println(unusedSlice, file)
}
