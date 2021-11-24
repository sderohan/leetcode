/*
Given a m x n matrix grid which is sorted in non-increasing order both row-wise and column-wise, return the number of negative numbers in grid.

Example 1:

Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
Output: 8
Explanation: There are 8 negatives number in the matrix.
Example 2:

Input: grid = [[3,2],[1,0]]
Output: 0
Example 3:

Input: grid = [[1,-1],[-1,-1]]
Output: 3
Example 4:

Input: grid = [[-1]]
Output: 1

Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 100
-100 <= grid[i][j] <= 100
*/

func countNegatives(grid [][]int) int {
	return countNeg(grid)
}

// O(n^2)
func bruteForce(grid [][]int) (counter int) {
	for _, row := range grid {
		for _, num := range row {
			if num < 0 {
				counter++
			}
		}
	}
	return
}

// O(n+m)
// Start traversing from last element of the first row
func countNeg(grid [][]int) (counter int) {
	row, col := len(grid), len(grid[0])
	i, j := 0, col-1
	for i < row && j > -1 {
		if grid[i][j] >= 0 {
			i++
		} else if grid[i][j] < 0 {
			counter += (row - i)
			j--
		}
	}
	return
}