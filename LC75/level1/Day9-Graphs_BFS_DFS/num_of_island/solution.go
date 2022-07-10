package main

//perform DFS
func dfs(grid [][]byte, row, col int) {
	nRow := len(grid)
	nCol := len(grid[0])

	//keeping bounds and whether we entered a visited cell
	if row < 0 || col < 0 || row >= nRow || col >= nCol || grid[row][col] == '0' {
		return
	}

	//mark as visited
	grid[row][col] = '0'

	//recursion : visiting neighbors
	dfs(grid, row-1, col)
	dfs(grid, row+1, col)
	dfs(grid, row, col-1)
	dfs(grid, row, col+1)
}

func numIslands(grid [][]byte) int {
	//if grid is nil or size 0 then return 0
	if len(grid) == 0 {
		return 0
	}

	//
	nRow := len(grid)
	nCol := len(grid[0])

	num_islands := 0

	//traverse through the grid
	for r := 0; r < nRow; r++ {
		for c := 0; c < nCol; c++ {

			//identify the DFS trigger
			if grid[r][c] == '1' {
				num_islands++
				dfs(grid, r, c)
			}
		}
	}

	return num_islands
}
