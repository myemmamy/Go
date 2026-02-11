// package palindromenumber
package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func main() {
	n := 123421
	m := 0
	fmt.Println(isPalindrome(n))
	fmt.Println(isPalindrome(m))
}
