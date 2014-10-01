/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package main

import "fmt"

var primes = make(map[int]bool)

func is_prime(num int) bool {
	if num % 2 == 0 {
		return false
	}
	for i := 3; i < (num/2)-1; i+=2 {
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
		for i := 3; i < max/2; i++ {
			if is_prime(i) {
				primes[i] = true
			}
		}
	}
}

func get_factor(num int) []int {
	factors := make([]int, 0)
	for prime := range primes {
		if num % prime == 0 {
			fmt.Println(prime)
			factors = append(factors, prime)
			if val := num/prime; val != 1 && val != 0{
				factors = append(factors, get_factor(num / prime)...)
			}
			return factors
		}
	}
	return []int{num}
}

func main() {
	num := 5000
	gen_primes(num)
	fmt.Println(primes)
	fmt.Println(get_factor(num))
}