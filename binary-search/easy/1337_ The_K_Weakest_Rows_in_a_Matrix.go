/*
1337. The K Weakest Rows in a Matrix
Easy

You are given an m x n binary matrix mat of 1's (representing soldiers) and 0's (representing civilians).
The soldiers are positioned in front of the civilians.
That is, all the 1's will appear to the left of all the 0's in each row.

A row i is weaker than a row j if one of the following is true:

The number of soldiers in row i is less than the number of soldiers in row j.
Both rows have the same number of soldiers and i < j.
Return the indices of the k weakest rows in the matrix ordered from weakest to strongest.

Example 1:

Input: mat =
[[1,1,0,0,0],
 [1,1,1,1,0],
 [1,0,0,0,0],
 [1,1,0,0,0],
 [1,1,1,1,1]],
k = 3
Output: [2,0,3]
Explanation:
The number of soldiers in each row is:
- Row 0: 2
- Row 1: 4
- Row 2: 1
- Row 3: 2
- Row 4: 5
The rows ordered from weakest to strongest are [2,0,3,1,4].
Example 2:

Input: mat =
[[1,0,0,0],
 [1,1,1,1],
 [1,0,0,0],
 [1,0,0,0]],
k = 2
Output: [0,2]
Explanation:
The number of soldiers in each row is:
- Row 0: 1
- Row 1: 4
- Row 2: 1
- Row 3: 1
The rows ordered from weakest to strongest are [0,2,3,1].


Constraints:

m == mat.length
n == mat[i].length
2 <= n, m <= 100
1 <= k <= m
matrix[i][j] is either 0 or 1.
*/

func kWeakestRows(mat [][]int, k int) []int {
	return countRows(mat, k)
}

// store the counts of 1 in each row along with row index
type node struct {
	index    int
	soldiers int
}

func countRows(mat [][]int, k int) []int {
	// declare the slice to hold the nodes
	counts := make([]node, len(mat), len(mat))

	// create the node with row index and count of the soldiers and store in array
	for idx, current_row := range mat {
		counts[idx] = node{index: idx, soldiers: lastIndex(current_row) + 1}
	}

	// sort the nodes based on the given condition
	sort.Slice(counts, func(i, j int) bool {
		si, sj := counts[i].soldiers, counts[j].soldiers
		return si == sj && counts[i].index < counts[j].index || si < sj
	})

	// declare array to store the final result
	result := make([]int, k, k)

	// get the result
	for i, row_num := range counts[:k] {
		result[i] = row_num.index
	}
	return result
}

// binary search to get the last occurence of 1 in given array
func lastIndex(arr []int) int {
	low, high := 0, len(arr)-1
	index := -1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == 1 {
			index = mid
			low++
		} else {
			high = mid - 1
		}
	}
	return index
}