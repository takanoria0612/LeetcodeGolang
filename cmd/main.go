package main

import (
	"GolangLeetcode/problems/ContainerWithMostWater"
	"GolangLeetcode/problems/TwoSum"

	"fmt"
)

func main() {
	// Container with most water
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result1 := ContainerWithMostWater.MaxArea(heights)
	fmt.Println("Solution 1 Max Area:", result1)
	result2 := ContainerWithMostWater.MaxArea2(heights)
	fmt.Println("Solution 2 Max Area:", result2)

	// Two Sum
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println("The indecies using solution1:", TwoSum.TwoSum(nums, target))  // 出力: [0, 1]
	fmt.Println("The indecies using solution2:", TwoSum.TwoSum2(nums, target)) // 出力: [0, 1]
}
