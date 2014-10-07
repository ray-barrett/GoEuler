/*


A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
package main

import (
	"fmt"
	"strconv"
)

func is_pallindrome(check string) bool {
	start, end := 0, len(check) - 1
	for start < end {
		if check[start] != check[end] {
			return false
		}	
		 start++ 
		 end-- 
	}
	return true
}

func main() {
	min, max := 100, 1000
	max_pal := 0
	for i:= min; i < max; i++ {
		for j := min; j < max; j++ {
			num := i * j
			if is_pallindrome(strconv.Itoa(num)) && num > max_pal {
				max_pal = num
			}
		}
	}
	fmt.Println(max_pal)
}