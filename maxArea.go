package main

import "math"

func maxArea(height []int) int {
	maxArea := 0
	l := 0
	r := len(height) - 1

	for l < r {
		h := float64(min(height[l], height[r]))
		maxArea = int(math.Max(float64(maxArea), float64(h*float64(r-l))))

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}
