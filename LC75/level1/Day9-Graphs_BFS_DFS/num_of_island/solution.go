package main

func numIslands(grid [][]byte) int {

	//if the grid is nil or size 0 return 0
	if grid == nil || len(grid) == 0 {
		return 0
	}

	numRow := len(grid)
	numCol := len(grid[0])
	num_island := 0
	r, c := 0, 0
	//traverse through rows
	for r < numRow {

		//traverse through col
		for c < numCol {
			num_island++
			grid[r][c] = '0' //mark as visited
			neighbors := []int{}
		}
	}
}
