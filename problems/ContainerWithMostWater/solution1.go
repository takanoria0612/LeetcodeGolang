package ContainerWithMostWater

import "math"

func MaxArea(height []int) int {
	max_area := 0
	l := 0
	r := len(height) - 1
	for l < r {
		h := float64(math.Min(float64(height[l]), float64(height[r])))
		max_area = int(math.Max(float64(max_area), h*float64(r-l)))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max_area
}
