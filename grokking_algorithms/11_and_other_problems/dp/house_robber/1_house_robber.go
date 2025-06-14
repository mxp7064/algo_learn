package main

import "fmt"

/*
=================================================================================
🧠 House Robber — Bottom-Up Dynamic Programming (O(n) time, O(n) space)
=================================================================================

This is LeetCode 198

Problem:
Given an array of non-negative integers representing the amount of money at each house,
return the maximum amount you can rob without robbing two adjacent houses.

Approach:
- For each house i, you decide:
    • Rob it: amount = houses[i] + dp[i-2]
    • Skip it: amount = dp[i-1]
- Transition: dp[i] = max(houses[i] + dp[i-2], dp[i-1])
- Base cases:
    dp[0] = houses[0]
    dp[1] = max(houses[0], houses[1])

🧪 Example Walkthrough: houses = [2, 7, 9, 3, 1]

We build the dp array where dp[i] represents the maximum amount that can be robbed from house 0 to i.

Step-by-step:

dp[0] = 2                     // Rob the first house: only one option
dp[1] = max(2, 7) = 7         // Choose between first (2) and second (7)
dp[2] = max(7, 2+9) = 11      // Rob first and third (2+9=11), better than just second (7)
dp[3] = max(11, 7+3) = 11     // Rob first and third (11) is better than second and fourth (7+3=10)
dp[4] = max(11, 11+1) = 12    // Rob first, third, and fifth (2+9+1 = 12)

Final dp array: [2, 7, 11, 11, 12]
Return dp[4] = 12

✅ Optimal houses to rob: house 0 (2), house 2 (9), house 4 (1)
Total = 2 + 9 + 1 = 12

🧠 Time and space complexity
Time Complexity: O(n)
Space Complexity: O(n) — we will optimize it to O(1) in next example: 2_house_robber_space_optimized.go
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

// getHousesToRob returns the max money that can be robbed without alerting the police (without robbing two adjacent houses)
func getHousesToRob(houses []int) int {
	n := len(houses)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return houses[0]
	}

	// dp[i] = max amount that can be robbed up to house i
	dp := make([]int, n)
	dp[0] = houses[0]
	dp[1] = maxInt(houses[0], houses[1])

	for i := 2; i < n; i++ {
		// Either rob this house (and skip the previous), or skip this house (chose previous)
		currentHouse := houses[i] + dp[i-2]
		previousHouse := dp[i-1]
		dp[i] = maxInt(currentHouse, previousHouse)
	}

	return dp[n-1] // The answer is the max we can rob up to the last house
}
