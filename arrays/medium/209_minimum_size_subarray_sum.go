/*
209. Minimum Size Subarray Sum
Medium

5167

161

Add to List

Share
Given an array of positive integers nums and a positive integer target, return the minimal length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr] of which the sum is greater than or equal to target. If there is no such subarray, return 0 instead.



Example 1:

Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.
Example 2:

Input: target = 4, nums = [1,4,4]
Output: 1
Example 3:

Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0


Constraints:

1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 105
*/

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	length := math.MaxInt64
	fp, sp := 0, 0
	temp_sum := 0
	for sp < n {
		temp_sum += nums[sp]
		for temp_sum >= target && fp < n {
			if length > (sp - fp + 1) {
				length = sp - fp + 1
			}
			temp_sum -= nums[fp]
			fp++
		}
		sp++
	}
	if length == math.MaxInt64 {
		length = 0
	}
	return length
}