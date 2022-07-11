package main

func climbStairs(n int) int {

	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)

	dp[1] = 1
	dp[2] = 2

	i := 3

	for i <= n {
		dp[i] = dp[i-1] + dp[i-2]
		i++
	}

	return dp[n]
}

func fibClimbStairs(n int) int {
	if n == 1 {
		return 1
	}

	first, second := 1, 2

	i := 3

	for i <= n {
		third := first + second
		first = second
		second = third
		i++
	}

	return second
}
