package main

func dfs(image [][]int, row, col, color, newColor int) {

	//set new color!
	if image[row][col] == color {
		image[row][col] = newColor

		if row >= 1 {
			dfs(image, row-1, col, color, newColor)
		}
		if col >= 1 {
			dfs(image, row, col-1, color, newColor)
		}
		if row+1 < len(image) {
			dfs(image, row+1, col, color, newColor)
		}
		if col+1 < len(image[0]) {
			dfs(image, row, col+1, color, newColor)
		}
	}

}

func floodFill(image [][]int, sr, sc, color int) [][]int {

	oldColor := image[sr][sc]

	//trigger DFS, if it's not the same color
	if oldColor != color {
		dfs(image, sr, sc, oldColor, color)
	}
	return image
}
