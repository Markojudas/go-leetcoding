package main

func uniquePaths(m int, n int) int {

	//Create a 2D array/slice d[m][n] = number of paths
	d := make([][]int, m)
	rows := make([]int, m*n)
	for i, startRow := 0, 0; i < m; i, startRow = i+1, startRow+n {
		endRow := startRow + n
		d[i] = rows[startRow:endRow:endRow]
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
