// https://leetcode.com/problems/encode-and-decode-strings/description/
//Design an algorithm to encode a list of strings to a single string. The encoded string is then decoded back to the original list of strings.
// Example 1:
// Input: ["lint","code","love","you"]
// Output: ["lint","code","love","you"]
// Explanation: One possible encode method is "lint:;code:;love:;you" where ":" is a delimiter and it can be used to decode back.
// Example 2:
// Input: ["we", "say", ":", "yes"]
// Output: ["we", "say", ":", "yes"]
// Explanation: One possible encode method is "we:;say:;::;yes" where ":" is a delimiter and it can be used to decode back.

package main

import (
	"fmt"
	"strings"
)

// Encodes a list of strings to a single string.
func encode(strs []string) string {
	// Use a delimiter that is not likely to appear in the strings
	delimiter := ":;"
	encoded := strings.Join(strs, delimiter)
	return encoded
}

// Decodes a single string to a list of strings.
func decode(s string) []string {
	// Use the same delimiter used in encoding
	delimiter := ":;"
	strs := strings.Split(s, delimiter)
	return strs
}

func main() {
	strs := []string{"lint", "code", "love", "you"}
	encoded := encode(strs)
	fmt.Println("Encoded:", encoded)

	decoded := decode(encoded)
	fmt.Println("Decoded:", decoded)
}
