/*
2062. Count Vowel Substrings of a String
Easy

268

89

Add to List

Share
A substring is a contiguous (non-empty) sequence of characters within a string.

A vowel substring is a substring that only consists of vowels ('a', 'e', 'i', 'o', and 'u') and has all five vowels present in it.

Given a string word, return the number of vowel substrings in word.



Example 1:

Input: word = "aeiouu"
Output: 2
Explanation: The vowel substrings of word are as follows (underlined):
- "aeiouu"
- "aeiouu"
Example 2:

Input: word = "unicornarihan"
Output: 0
Explanation: Not all 5 vowels are present, so there are no vowel substrings.
Example 3:

Input: word = "cuaieuouac"
Output: 7
Explanation: The vowel substrings of word are as follows (underlined):
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"


Constraints:

1 <= word.length <= 100
word consists of lowercase English letters only.
*/

func countVowelSubstrings(word string) int {
	n := len(word)
	mp := map[byte]struct{}{
		'a': struct{}{},
		'e': struct{}{},
		'i': struct{}{},
		'o': struct{}{},
		'u': struct{}{},
	}
	vowelStrings := 0
	for i := 0; i < n; i++ {
		temp := make(map[byte]struct{}, 5)
		for j := i; j < n; j++ {
			if _, ok := mp[word[j]]; ok {
				temp[word[j]] = struct{}{}
				if len(temp) == 5 {
					vowelStrings++
				}
			} else {
				break
			}
		}
	}
	return vowelStrings
}