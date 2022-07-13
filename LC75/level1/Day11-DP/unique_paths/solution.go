package main

func uniquePaths(m int, n int) int {

	//Create a 2D array/slice d[m][n] = number of paths
	d := make([][]int, m)
	for idx := range d {
		d[idx] = make([]int, n)
	}

	//initiate all cells to 1
	for idx, row := range d {
		for idx2 := range row {
			d[idx][idx2] = 1
		}
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			d[row][col] = d[row-1][col] + d[row][col-1]
		}
	}

	return d[m-1][n-1]
}
