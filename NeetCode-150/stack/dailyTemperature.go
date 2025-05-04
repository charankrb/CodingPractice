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

import (
	"fmt"
	"reflect"
)

func dailyTemperatures(temp []int) []int {
	// Initialize the result array with the same length as the input array, filled with zeros
	result := make([]int, len(temp))
	// Initialize an empty stack to store indices of temperatures
	stack := []int{}
	// Iterate through the temperatures array
	for i, t := range temp {
		// While the stack is not empty and the current temperature is greater than the temperature
		// at the index stored at the top of the stack
		for len(stack) > 0 && t > temp[stack[len(stack)-1]] {
			// Get the index of the previous day with a lower temperature
			index := stack[len(stack)-1]

			// Remove the top element from the stack as the current temperature is warmer
			stack = stack[:len(stack)-1]

			// Calculate the number of days until a warmer temperature and store it in the result array
			result[index] = i - index
		}
		// Push the current index onto the stack
		stack = append(stack, i)
	}
	return result
}

func main() {
	// Test cases
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}}, // General case
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},                             // Increasing temperatures
		{[]int{30, 20, 10, 5}, []int{0, 0, 0, 0}},                              // Decreasing temperatures
		{[]int{30, 30, 30, 30}, []int{0, 0, 0, 0}},                             // Constant temperatures
		{[]int{30}, []int{0}},                                                  // Single temperature
		{[]int{}, []int{}},                                                     // Empty input
		{[]int{50, 60, 50, 60, 50, 60}, []int{1, 0, 1, 0, 1, 0}},               // Alternating temperatures
		{[]int{70, 71, 72, 73, 74, 75, 76, 77}, []int{1, 1, 1, 1, 1, 1, 1, 0}}, // Strictly increasing temperatures
		{[]int{80, 70, 60, 50, 40, 30}, []int{0, 0, 0, 0, 0, 0}},               // Strictly decreasing temperatures
	}

	// Run the test cases
	for i, testCase := range testCases {
		result := dailyTemperatures(testCase.input)
		fmt.Printf("Test Case %d: Input: %v\n", i+1, testCase.input)
		fmt.Printf("Expected: %v\n", testCase.expected)
		fmt.Printf("Got: %v\n", result)

		// Validate the result
		if reflect.DeepEqual(result, testCase.expected) {
			fmt.Printf("Test Case %d Passed\n\n", i+1)
		} else {
			fmt.Printf("Test Case %d Failed\n\n", i+1)
		}
	}
}
