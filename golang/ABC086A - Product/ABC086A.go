package main

import (
	"fmt"
)

func main() {
	var a, b int
	var c string
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	if (a*b)%2 == 0 {
		c = "Even"
	} else {
		c = "Odd"
	}

	fmt.Printf("%s", c)
}
