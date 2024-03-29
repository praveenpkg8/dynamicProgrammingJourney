/*
Link: https://leetcode.com/problems/n-th-tribonacci-number/

1137. N-th Tribonacci Number
Easy

1517

95

Add to List

Share
The Tribonacci sequence Tn is defined as follows: 

T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.

Given n, return the value of Tn.

 

Example 1:

Input: n = 4
Output: 4
Explanation:
T_3 = 0 + 1 + 1 = 2
T_4 = 1 + 1 + 2 = 4
Example 2:

Input: n = 25
Output: 1389537
 

Constraints:

0 <= n <= 37
The answer is guaranteed to fit within a 32-bit integer, ie. answer <= 2^31 - 1.
*/

package main

func tribonacci(n int) int {
    if n < 2 {
        return n
    }
    trib, prev, prev2 := 1, 1, 0
    for n > 2 {
        trib, prev, prev2 = trib + prev + prev2, trib, prev
        n--
    }
    return trib
}

func tribonacciMemmo(n int) int {
    if (n == 0) {
        return 0
    } else if (n <= 2) {
        return 1
    }
    
    memCache := make([]int, n + 1)
    memCache[0], memCache[1], memCache[2] = 0, 1, 1
    for i := 3; i <= n; i += 1 {
        memCache[i] = memCache[i - 3] + memCache[i - 2] + memCache[i - 1]
    }
    return memCache[n]
    
}