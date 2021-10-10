package easy

import "fmt"

// link problem:
// https://leetcode.com/problems/island-perimeter/

func IslandPerimeter(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res += 4
				if i-1 >= 0 && grid[i-1][j] == 1 {
					res--
				}
				if j-1 >= 0 && grid[i][j-1] == 1 {
					res--
				}
				if i+1 < m && grid[i+1][j] == 1 {
					res--
				}
				if j+1 < n && grid[i][j+1] == 1 {
					res--
				}
			}
		}
	}
	return res
}

func Main_IslandPerimeter() {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	fmt.Println("=========IslandPerimeter=========")
	fmt.Println("Input\t: ", grid)
	fmt.Println("Output\t: ", IslandPerimeter(grid))
	fmt.Println("")
}
