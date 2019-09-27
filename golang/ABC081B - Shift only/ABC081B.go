package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var a int

	fmt.Scan(&a)

	numarr := strings.Split(a, " ")

	var result int = 0

	for {
		count := 0
		for _, r := range numarr {
			num, _ := strconv.Atoi(r)
			if (num%2 == 0) && (num != 0) {
				count++
			}
		}
		if len(numarr) == count {
			result++
		} else {
			break
		}
	}
	fmt.Printf("%d", result)
}
