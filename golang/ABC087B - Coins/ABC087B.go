package main

import(
	"fmt"
)

func main() {
	var ans int
	var fiveHandredYen int
	var HandredYen int
	var FiftyYen int
	var total int
	fmt.Scanf("%d", &fiveHandredYen)
	fmt.Scanf("%d", &HandredYen)
	fmt.Scanf("%d", &FiftyYen)
	fmt.Scanf("%d", &total)

	for h := 0; h <= fiveHandredYen; h++ {
		for i := 0; i <= HandredYen; i++ {
			for j := 0; j <= FiftyYen; j++ {
				if total == (h * 500 + i * 100 + j * 50){
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}
