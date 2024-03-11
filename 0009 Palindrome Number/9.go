package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x <= 0 {
		return false
	}

	if x < 10 {
		return true
	}

	strInt := strconv.Itoa(x)
	for i := 0; i <= int(float64(len((strInt))/2)); i++ {
		if string(strInt[i]) != string(strInt[(len(strInt)-1)-i]) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(10))

}
