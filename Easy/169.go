/*
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.



Example 1:

Input: nums = [3,2,3]
Output: 3

Example 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2

*/

package main

func majorityElement(nums []int) int {
	count, res := 0, 0

	for _, num := range nums {
		if count == 0 {
			res = num
		}

		if res == num {
			count++
		} else {
			count--
		}
	}

	return res
}

/* APPROACH
We use the boyre moor algorithm

We keep two flag, counter and the current large number,

whenever we get a number that is same as the current large number, we do count++

else we do count --

whenever w get cont = 0, we will assume the majority number is the current number
*/
