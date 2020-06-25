package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(b)
	for a != b {
		if b < a {
			b++
		} else {
			b--
		}
		fmt.Print(" ", b)
	}
	fmt.Println()
}
