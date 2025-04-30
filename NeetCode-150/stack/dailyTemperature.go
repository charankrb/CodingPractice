//https://leetcode.com/problems/daily-temperatures/description/
// Daily Temperatures
//Given an array of integers temperatures represents the daily temperatures, return an array answer such that
// answer[i] is the number of days you have to wait after the ith day to get a warmer temperature.
// If there is no future day for which this is possible, keep answer[i] == 0 instead.
// Example 1:
// Input: temperatures = [73,74,75,71,69,72,76,73]
// Output: [1,1,4,2,1,1,0,0]
// Explanation:
// At index 0, the next warmer temperature is 74 (1 day away).
// At index 1, the next warmer temperature is 75 (1 day away).
// At index 2, there is no future day for which this is possible, so keep answer[2] == 0.
// At index 3, the next warmer temperature is 72 (1 day away).
// At index 4, the next warmer temperature is 76 (1 day away).
// At index 5, the next warmer temperature is 76 (1 day away).
// At index 6, there is no future day for which this is possible, so keep answer[6] == 0.
// At index 7, there is no future day for which this is possible, so keep answer[7] == 0.

// Example 2:
// Input: temperatures = [30,40,50,60]
// Output: [1,1,1,0]

package main
