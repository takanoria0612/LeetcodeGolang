package main

import (
	"GolangLeetcode/problems/ContainerWithMostWater"
	"fmt"
)

func main() {
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result1 := ContainerWithMostWater.MaxArea(heights) // MaxAreaは公開関数（大文字で始まる）
	fmt.Println("Solution 1 Max Area:", result1)
}
