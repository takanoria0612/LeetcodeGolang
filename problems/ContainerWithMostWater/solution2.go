package ContainerWithMostWater

func MaxArea2(height []int) int {
	retVal := 0
	for i := 0; i < len(height); i++ {
		localMax := 0
		for j := len(height) - 1; j >= i; j-- {

			//　ここである基準を設けて、不要な計算を避けてるんだよね。
			//　だから別にmaxContainerPossible := (j - i) * height[j]でもいいと思う
			maxContainerPossible := (j - i) * height[i]
			if maxContainerPossible < retVal || maxContainerPossible < localMax {
				break
			}
			containerSize := (j - i) * min(height[i], height[j])
			localMax = max(localMax, containerSize)
		}
		retVal = max(retVal, localMax)
	}
	return retVal
}

// 最小値を返す関数
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 最大値を返す関数
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
