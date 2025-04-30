// https://leetcode.com/problems/min-stack/description/
//Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
//Implement the MinStack class:
//MinStack() initializes the stack object.
//void push(int val) pushes the element val onto the stack.
//void pop() removes the element on the top of the stack.
//int top() gets the top element of the stack.
//int getMin() retrieves the minimum element in the stack.
//You must implement a solution with O(1) time complexity for each function.

//Example 1:
//Input : ["MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"]
//Output : [null, null, null, null, -2, null, 0, -1]
//Explanation : perform the following operations:
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin(); // return -3
//minStack.pop();
//minStack.top(); // return 0
//minStack.getMin(); // return -2

package main

type MinStack struct {
	stack []int // Stack to store the elements
	min   []int // Stack to store the minimum elements
}

func constructMinStack() MinStack {
	return MinStack{ // Constructor to initialize the MinStack
		stack: []int{},
		min:   []int{},
	}
}

func (ms *MinStack) Push(val int) { // Push an element onto the stack, complexity O(1)
	ms.stack = append(ms.stack, val)                      // Add the value to the main stack
	if len(ms.min) == 0 || val <= ms.min[len(ms.min)-1] { // If the min stack is empty or the new value is less than or equal to the current minimum
		ms.min = append(ms.min, val) // Add the value to the min stack
	}
}

func (ms *MinStack) Pop() { // Pop the top element from the stack, Complexity O(1)
	if len(ms.stack) == 0 { // If the stack is empty, do nothing
		return
	}
	val := ms.stack[len(ms.stack)-1]      // Get the top value from the main stack
	ms.stack = ms.stack[:len(ms.stack)-1] // Remove the top value from the main stack
	if val == ms.min[len(ms.min)-1] {     // If the popped value is equal to the current minimum
		ms.min = ms.min[:len(ms.min)-1] // Remove it from the min stack as well
	}
}

func (ms *MinStack) GetMin() int { // Get the minimum element in the stack, Complexity O(1)
	if len(ms.min) == 0 { // If the min stack is empty, return -1 (or any other error value)
		return -1
	}
	return ms.min[len(ms.min)-1] // Return the top value from the min stack
}

func (ms *MinStack) Top() int { // Get the top element of the stack, Complexity O(1)
	if len(ms.stack) == 0 { // If the stack is empty, return -1 (or any other error value)
		return -1
	}
	return ms.stack[len(ms.stack)-1] // Return the top value from the main stack
}

func main() {
	// Example usage of MinStack
	minStack := constructMinStack()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	println(minStack.GetMin()) // return -3
	minStack.Pop()
	println(minStack.Top())    // return 0
	println(minStack.GetMin()) // return -2
}
