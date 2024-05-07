package TwoSum

func TwoSum2(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		result := target - num
		if _, ok := numMap[result]; ok {
			return []int{numMap[result], i}
		}
		numMap[num] = i
	}
	return nil
}
