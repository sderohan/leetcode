/*
https://leetcode.com/problems/uncommon-words-from-two-sentences/
884. Uncommon Words from Two Sentences
Easy

830

128

Add to List

Share
A sentence is a string of single-space separated words where each word consists only of lowercase letters.

A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.

Given two sentences s1 and s2, return a list of all the uncommon words. You may return the answer in any order.



Example 1:

Input: s1 = "this apple is sweet", s2 = "this apple is sour"
Output: ["sweet","sour"]
Example 2:

Input: s1 = "apple apple", s2 = "banana"
Output: ["banana"]


Constraints:

1 <= s1.length, s2.length <= 200
s1 and s2 consist of lowercase English letters and spaces.
s1 and s2 do not have leading or trailing spaces.
All the words in s1 and s2 are separated by a single space.
*/

func uncommonFromSentences(s1 string, s2 string) []string {
	map_s1, map_s2 := getMap(s1), getMap(s2)
	var result []string
	result = getUnique(map_s1, map_s2)
	result = append(result, getUnique(map_s2, map_s1)...)
	return result
}

func getMap(s string) map[string]int {
	n := len(s)
	map_s := make(map[string]int)
	temp := ""
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			map_s[temp]++
			temp = ""
		} else {
			temp += string(s[i])
		}
	}
	// for the last word
	map_s[temp]++
	return map_s
}

func getUnique(src, dest map[string]int) []string {
	var result []string
	for key, value := range src {
		if value == 1 {
			if _, ok := dest[key]; !ok {
				result = append(result, key)
			}
		}
	}
	return result
}