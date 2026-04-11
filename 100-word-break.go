//go:build ignore
// +build ignore

/*
Name: Word Break
Source: Leetcode
Level: Medium
TC: ??
Desc: Given a string s and a dictionary of strings wordDict, return true if s can be
segmented into a space-separated sequence of one or more
dictionary words.
Note: The same word in the dictionary may be reused multiple times in the segmentation.
Link: https://leetcode.com/problems/word-break
PD: Note many leetcode have some trick, that all words in the dictionary must match one or more times from the beggining to the end of the string
*/

package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func wordBreak2(s string, wordDict []string) bool {
	l := len(s)
	for offset := 0; offset < l; {
		anyFound := false
		for _, word := range wordDict {
			if strings.HasPrefix(s[offset:], word) {
				offset = offset + len(word)
				anyFound = true
			}
		}
		if !anyFound {
			return false
		}
	}
	return true
}

// searchSeq
// given a s and a list of words, returns true if any of the words can be
// contactenated and produce s, exact match
func searchSeq(s string, wordDict []string) bool {
	index := 0
	l := len(s)
	// any of the words should exists
	for index = 0; index < l; {
		// search any word
		found := false
		for _, word := range wordDict {
			if strings.HasPrefix(s[index:], word) {
				index += len(word)
				found = true
			}
		}
		// at least one found?
		if !found {
			return false
		}
	}
	// me sure not characters are left
	return index == len(s)
}

func wordBreak(s string, wordDict []string) bool {
	// sort wordDict by len first
	sort.Slice(wordDict, func(i, j int) bool {
		return len(wordDict[i]) < len(wordDict[j])
	})

	for len(wordDict) > 0 {
		if searchSeq(s, wordDict) {
			return true
		}
		wordDict = wordDict[1:len(wordDict)]
	}
	return false
}

func pre2(s string, wordDict []string) bool {
	var dictionary []string
	for _, word := range wordDict {
		dictionary = append(dictionary, word)
	}
	for _, word := range dictionary {
		if !strings.Contains(s, word) {
			return false
		}
	}
	return true
}

func pre(s string, wordDict []string) bool {
	starts := false
	ends := false
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			starts = true
		}
		if strings.HasSuffix(s, word) {
			ends = true
		}
	}
	return starts && ends
}

func searchR(s string, index int, wordDict []string) bool {
	if !pre(s, wordDict) && !pre2(s, wordDict) {
		return false
	}

	for _, word := range wordDict {
		if strings.HasPrefix(s[index:], word) {
			next := index + len(word)
			if next == len(s) {
				// one combination was found
				return true
			}
			if next < len(s) {
				if searchR(s, next, wordDict) {
					// word with a next combiation was found
					return true
				}
			}
			// here we dont know yet, maybe another word
		}
	}
	// after looking at all combinations not found
	return false
}

func wordBreakExp(s string, wordDict []string) bool {
	rexp := "^("
	l := len(wordDict)
	for i, word := range wordDict {
		if i < l-1 {
			rexp += fmt.Sprintf("%v|", word)
		} else {
			rexp += fmt.Sprintf("%v", word)
		}
	}
	rexp += ")+$"
	fmt.Println(rexp)
	r, err := regexp.Compile(rexp)
	if err != nil {
		fmt.Println("err compiling:", err)
		return false
	}
	match := r.MatchString(s)
	return match

}

func wordBreakV2(s string, wordDict []string) bool {
	return searchR(s, 0, wordDict)
}

func main() {
	var testCases = []struct {
		s        string
		wordDict []string
		expected bool
	}{
		{s: "leetcode", wordDict: []string{"leet", "code"}, expected: true},
		{s: "applepenapple", wordDict: []string{"apple", "pen"}, expected: true},
		{s: "catsandog", wordDict: []string{"cats", "dog", "sand", "and", "cat"}, expected: false},
		{s: "bb", wordDict: []string{"a", "b", "bbb", "bbbb"}, expected: true}, // the word 'b' can be reused
		{s: "cars", wordDict: []string{"car", "ca", "rs"}, expected: true},
		{s: "catsandogcat", wordDict: []string{"cats", "dog", "sand", "and", "cat", "an"}, expected: true},
		{s: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}, expected: false},
		{s: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}, expected: false},
	}
	for i, t := range testCases {
		val := wordBreakExp(t.s, t.wordDict)
		if val != t.expected {
			fmt.Printf("case %v wrong: expected %v\n", i, t.expected)
		}
	}
}

// "catsandogcat"
// ["cats","dog","sand","and","cat","an"]
