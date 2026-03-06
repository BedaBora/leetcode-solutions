/*
Candy

There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.

You are giving candies to these children subjected to the following requirements:

    Each child must have at least one candy.
    Children with a higher rating get more candies than their neighbors.

Return the minimum number of candies you need to have to distribute the candies to the children.



Example 1:

Input: ratings = [1,0,2]
Output: 5
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.

Example 2:

Input: ratings = [1,2,2]
Output: 4
Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
The third child gets 1 candy because it satisfies the above two conditions.

*/

package main

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	ret, up, down, peak := 1, 0, 0, 0

	for i := range n - 1 {
		prev, curr := ratings[i], ratings[i+1]

		if prev < curr {
			up++
			down = 0
			peak = up
			ret += 1 + up
		} else if prev == curr {
			up, down, peak = 0, 0, 0
			ret += 1
		} else {
			up = 0
			down++
			ret += 1 + down
			if peak >= down {
				ret--
			}
		}
	}

	return ret
}

/* APPROACH
The essence of the one-pass greedy algorithm lies in these three variables: Up, Down, and Peak. They serve as counters for the following:

Up: Counts how many children have increasing ratings from the last child. This helps us determine how many candies we need for a child with a higher rating than the previous child.

Down: Counts how many children have decreasing ratings from the last child. This helps us determine how many candies we need for a child with a lower rating than the previous child.

Peak: Keeps track of the last highest point in an increasing sequence. When we have a decreasing sequence after the peak, we can refer back to the Peak to adjust the number of candies if needed.

*/
