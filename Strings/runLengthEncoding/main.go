package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "AAAAAAAAAAAAABBCCCCDD1111"
	actual := RunLengthEncoding(input)
	fmt.Println(actual)
}

func RunLengthEncoding(str string) string {
	encodedString := []byte{}
	runLength := 1

	for i := 1; i < len(str); i++ {
		currentChar := str[i]
		previousChar := str[i-1]
		if (currentChar != previousChar) || runLength == 9 {
			//fmt.Println(strconv.Itoa(runLength)[0])
			encodedString = append(encodedString, strconv.Itoa(runLength)[0])
			encodedString = append(encodedString, previousChar)
			runLength = 0
		}
		runLength++
	}

	encodedString = append(encodedString, strconv.Itoa(runLength)[0])
	encodedString = append(encodedString, str[len(str)-1])
	return string(encodedString)
}
