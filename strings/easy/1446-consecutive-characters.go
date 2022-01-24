/*
1446. Consecutive Characters
Easy

1143

22

Add to List

Share
The power of the string is the maximum length of a non-empty substring that contains only one unique character.

Given a string s, return the power of s.



Example 1:

Input: s = "leetcode"
Output: 2
Explanation: The substring "ee" is of length 2 with the character 'e' only.
Example 2:

Input: s = "abbcccddddeeeeedcba"
Output: 5
Explanation: The substring "eeeee" is of length 5 with the character 'e' only.


Constraints:

1 <= s.length <= 500
s consists of only lowercase English letters.
*/

func maxPower(s string) int {
	n := len(s)
	maxCount := 1
	for i := 1; i < n; i++ {
		count := 1
		// look for consecutive char
		for i < n && s[i-1] == s[i] {
			count++
			i++
		}
		// update the maximum count
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}