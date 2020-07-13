package util

import (
	"fmt"

	"../utildepth2"
)

func LenWrapperU(ss []string) int {
	return utildepth2.LenWrapper(ss)
}

func NotUsed() {
	b := []string{"just", "something"}
	a := len(b)
	for i, v := range b {
		if i == a-1 {
			fmt.Println("Last element", v)
			continue
		}
		fmt.Println("Element", v)
	}
}
