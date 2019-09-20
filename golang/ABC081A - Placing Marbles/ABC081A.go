package main

import (
	"fmt"
	"strings"
)

func main() {
	var result int
	var masu string
	fmt.Scanf("%s", &masu)
	arr := strings.Split(masu, "")
	for _, r := range arr {
		if r == "1" {
			result++
		}
	}
	fmt.Printf("%d", result)
}
