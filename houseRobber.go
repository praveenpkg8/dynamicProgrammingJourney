/*
Link: https://leetcode.com/problems/house-robber/

198. House Robber
Medium

10959

241

Add to List

Share
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

 

Example 1:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.
 

Constraints:

1 <= nums.length <= 100
0 <= nums[i] <= 400
*/
package main

func rob(nums []int) int {
    houseLen := len(nums)
    if houseLen == 0 {
        return 0
    }
    if houseLen == 1 {
        return nums[0]
    }
    dp := make([]int, houseLen)
    dp[0] = nums[0]
    dp[1] = findMax(nums[0], nums[1])
    
    for i := 2; i < houseLen; i ++ {
        dp[i] = findMax(nums[i] + dp[i-2], dp[i-1])
    }
    return dp[houseLen-1]
    
}

func findMax(x, y int) int {
    if x >= y {
        return x
    }
    return y
}
