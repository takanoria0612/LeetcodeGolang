package main

import (
	"GolangLeetcode/problems/ContainerWithMostWater"
	"fmt"
)

func main() {
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result1 := ContainerWithMostWater.MaxArea(heights)
	fmt.Println("Solution 1 Max Area:", result1)
	result2 := ContainerWithMostWater.MaxArea2(heights)
	fmt.Println("Solution 2 Max Area:", result2)
}
