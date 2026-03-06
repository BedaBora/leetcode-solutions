/*
Trapping Rain Water

Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

*/

package main

func trap(height []int) int {
	totalWater := 0
	n := len(height)
	if n == 0 {
		return 0
	}

	left, right := 0, n-1
	maxLeft, maxRight := height[left], height[right]

	for left < right {
		if maxLeft < maxRight {
			left++
			maxLeft = max(maxLeft, height[left])
			totalWater += maxLeft - height[left]
		} else {
			right--
			maxRight = max(maxRight, height[right])
			totalWater += maxRight - height[right]
		}
	}
	return totalWater
}

/* APPROACH

we take 2 pointer to keep track of the largest height on the left and right

then add up the water level - height given (because we cannot hold water at that level)

*/
