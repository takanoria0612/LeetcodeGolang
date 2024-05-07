package main

import (
	"GolangLeetcode/problems/ContainerWithMostWater"
	"GolangLeetcode/problems/GroupAnagrams"
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

	// Group Anagrams
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result_G1 := GroupAnagrams.GroupAnagrams(strs)
	fmt.Println("Group Anagrams using solution1:", result_G1) // Output: [[eat tea ate] [bat] [tan nat]]

	result_G2 := GroupAnagrams.GroupAnagrams2(strs)
	fmt.Println("Group Anagrams using solution2:", result_G2) // Output: [[eat tea ate] [bat] [tan nat]]

	result_G3 := GroupAnagrams.GroupAnagrams3(strs)
	fmt.Println("Group Anagrams using solution3:", result_G3) // Output: [[eat tea ate] [bat] [tan nat]]
}
