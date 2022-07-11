package main

//using Memoization
// storing newly found fib numbers to prevent recalculations
//starting with 0 and 1 as these are known.

var cache = map[int]int{
	0: 0,
	1: 1,
}

func fib(n int) int {

	//first check if the fib number is cached
	if _, ok := cache[n]; ok {
		return cache[n]
	}

	//cache the fib num
	cache[n] = fib(n-1) + fib(n-2)

	//return the fib num
	return cache[n]
}

//not using a map/dictionary to save memory
//Iterative approach keeping the two previous numbers

func fib2(n int) int {

	if n <= 1 {
		return n
	}

	current, prev1, prev2 := 0, 1, 0

	for i := 2; i <= n; i++ {
		current = prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return current
}
