//https://leetcode.com/problems/car-fleet/description/
// Car Fleet
// There are n cars going to the same destination along a one-lane road. The destination is target miles away.
// Each car i has a constant speed of speed[i] miles per hour (0 <= i < n).
// Note that the other car speeds are not given, but we can assume that all cars are going to the same destination at the same time.
//
// A car fleet is some non-empty set of cars driving at the same speed, such that no car in the fleet overtakes another car in the fleet.
// Note that the car fleet is not necessarily contiguous (i.e., there can be other cars in between).

// Return the total number of car fleets that will arrive at the destination.
//Example 1:
// Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
// Output: 3
// Explanation:
// The car fleets are [10,8], [0,5,3], and [2].
// The car fleet [10,8] will arrive at the destination at the same time, so they form a fleet.
// The car fleet [0,5,3] will arrive at the destination at the same time, so they form a fleet.
// The car fleet [2] will arrive at the destination at the same time, so it forms a fleet by itself.

package main

import (
	"fmt"
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	l := len(position) // Get the number of cars

	// Create a slice to store pairs of position and speed for each car
	carMap := make([][2]int, l)
	for i := 0; i < l; i++ {
		carMap[i] = [2]int{position[i], speed[i]} // Store position and speed as a pair
	}

	// Sort the cars by their starting position in descending order
	sort.Slice(carMap, func(i, j int) bool {
		return carMap[i][0] > carMap[j][0] // Cars closer to the target come first
	})
	// After sorting, carMap will look like: {{10,2},{8,4},{5,1},{3,3},{0,1}}

	// Initialize a stack to store the time each car takes to reach the target
	stack := []float64{}

	// Iterate through the sorted cars
	for _, m := range carMap {
		// Calculate the time for the current car to reach the target
		time := float64(target-m[0]) / float64(m[1]) // Time = (distance to target) / speed

		// Add the time to the stack
		stack = append(stack, time)

		// If the current car's time is less than or equal to the previous car's time,
		// it means the current car will join the same fleet as the previous car
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1] // Remove the current car's time from the stack
		}
	}

	// The number of fleets is equal to the size of the stack
	return len(stack)
}

func main() {
	// Test cases
	testCases := []struct {
		target   int
		position []int
		speed    []int
		expected int
	}{
		{12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}, 3},   // General case
		{10, []int{3}, []int{3}, 1},                            // Single car
		{10, []int{}, []int{}, 0},                              // No cars
		{15, []int{10, 8, 5}, []int{2, 4, 1}, 2},               // All cars form separate fleets
		{20, []int{10, 8, 5}, []int{4, 4, 4}, 3},               // All cars form one fleet
		{100, []int{90, 80, 70, 60}, []int{10, 20, 30, 40}, 1}, // Strictly decreasing speeds
		{30, []int{5, 10, 15, 20}, []int{1, 1, 1, 1}, 4},       // All cars have the same speed
	}

	// Run the test cases
	for i, testCase := range testCases {
		result := carFleet(testCase.target, testCase.position, testCase.speed)
		fmt.Printf("Test Case %d: Target: %d, Position: %v, Speed: %v\n", i+1, testCase.target, testCase.position, testCase.speed)
		fmt.Printf("Expected: %d, Got: %d\n", testCase.expected, result)

		// Validate the result
		if result == testCase.expected {
			fmt.Printf("Test Case %d Passed\n\n", i+1)
		} else {
			fmt.Printf("Test Case %d Failed\n\n", i+1)
		}
	}
}
