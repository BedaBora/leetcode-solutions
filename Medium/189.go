/*
Rotate Array

Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

Example 1:

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

Example 2:

Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
*/

package main

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	rotate := k % n

	if rotate == 0 {
		return
	}

	reverse(nums, 0, n-1)
	reverse(nums, 0, rotate-1)
	reverse(nums, rotate, n-1)
}

/* We first reverse the array, then dso modulo of n%k for the total no of rotates
since n rotation will bring the array to its former state

then we reverse it upto k
and then reverse if from k+1 to n
*/
