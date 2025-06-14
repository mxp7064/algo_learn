package main

import "fmt"

/*
===================================================================================
🧠 House Robber — Space-Optimized Dynamic Programming (O(n) Time, O(1) Space)
===================================================================================

Problem:
You are given an array of non-negative integers representing the amount of money
in each house. You cannot rob two adjacent houses.

Goal:
Return the maximum amount of money you can rob without alerting the police.

Approach:
- This is a space-optimized version of bottom-up dynamic programming.
- Instead of using a full dp array, we use two variables:
    prev1 → max amount up to previous house
    prev2 → max amount up to the house before that
- At each step:
    current = max(houses[i] + prev2, prev1)
- Then shift values forward: prev2 ← prev1, prev1 ← current

This is similar to the Fibonacci space-optimized pattern.

Time Complexity:  O(n)
Space Complexity: O(1)

🧪 Example Walkthrough: houses = [2, 7, 9, 3, 1]
Step-by-step:

dp[0] = 2         // Rob the first house
dp[1] = max(2, 7) = 7
dp[2] = max(7, 2+9) = 11
dp[3] = max(11, 7+3) = 11
dp[4] = max(11, 11+1) = 12

✅ Final Result: 12 → rob house 0 (2), house 2 (9), house 4 (1)

*/

func main() {
	fmt.Println(getHousesToRob([]int{2, 7, 9, 3, 1})) // Output: 12
	fmt.Println(getHousesToRob([]int{1, 2, 3, 1}))    // Output: 4
	fmt.Println(getHousesToRob([]int{0}))             // Output: 0
	fmt.Println(getHousesToRob([]int{5}))             // Output: 5
}

// maxInt returns the greater of a and b
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// getHousesToRob returns the max amount that can be robbed without robbing two adjacent houses
func getHousesToRob(houses []int) int {
	n := len(houses)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return houses[0]
	}

	// Initialize values for the first two houses
	prev2 := houses[0]                    // Max up to house i-2
	prev1 := maxInt(houses[0], houses[1]) // Max up to house i-1
	current := prev1                      // Max up to current

	for i := 2; i < n; i++ {
		// Either rob current house (and add to prev2) or skip it (take prev1)
		current = maxInt(houses[i]+prev2, prev1)

		// Shift values for the next iteration
		prev2 = prev1
		prev1 = current
	}

	return current
}
