/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package main

import "fmt"
import "sort"

type sortableArray []int
// Methods required by sort.Interface.
func (s sortableArray) Len() int {
    return len(s)
}
// "Less" is actually greater than because I want decending order sort.
func (s sortableArray) Less(i, j int) bool {
    return s[i] > s[j]
}
func (s sortableArray) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

var primes = make(map[int]bool)

func is_prime(num int) bool {
	if num % 2 == 0 {
		return false
	}
	for i := 3; i < (num/i)-1; i+=2 {
		if num % i == 0 {
			return false
		} 
	}
	return true
}

func gen_primes(max int) {
	if max == 2 {
		primes[2] = true
	} else {
		for i := 3; i < max/i-1; i += 2 {
			if is_prime(i) {
				primes[i] = true
			}
		}
	}
}

func get_factor(num int) sortableArray {
	factors := make(sortableArray, 0)
	for prime := range primes {
		if num % prime == 0 {
			factors = append(factors, prime)
			if val := num/prime; val != 1 && val != 0{
				factors = append(factors, get_factor(num / prime)...)
			}
			return factors
		}
	}
	factors = append(factors, num)
	return factors
}

func main() {
	num := 600851475143
	// Generate primes to factor with
	gen_primes(num)
	// Get sortableArry of factors
	factors := get_factor(num)
	// Sort
	sort.Sort(factors)
	// ??????
	// Profit
	fmt.Println(factors[0])
}