package main

import(
	"fmt"
)

func main() {
	for cock := 0; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++{
			for chicken := 0; chicken <= 100; chicken += 3{
				if (cock * 5 + hen * 3 + chicken / 3 == 100) && cock + hen + chicken == 100 {
					fmt.Printf("鸡翁 %d 只，母鸡 %d , 小鸡 %d \n", cock, hen, chicken)
				}
			}
		}
	}
}
