package leetcode0014

import "strings"

/*
Simple solution LCP(S1, S2, S3, ..., Sn) can be represented as
LCP(LCP(LCP(S1, S2), S3), ..., Sn)
Time complexity: O(S), S is sum of all characters in all strings
Space complexity: O(1), we use constant extra space.
*/
func longestCommonPrefixV0(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for _, s := range strs[1:] {
		for strings.Index(s, prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}

	return prefix
}

/*
We scan the string vertically. We compare characters from top to bottom on the
same column (same character index of the string) before moving on to the next
column.

Time complexity: O(S), S is sum of all characters in all strings
Space complexity: O(1), we use constant extra space.
*/

func longestCommonPrefixV1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := range strs[0] {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if (len(strs[j]) == i) || (strs[j][i] != c) {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

/*
Solution based on a trie
Time complexity: O(S), S is sum of all characters in all strings to build trie.
                 O(m), m is the length of common prefix
Space complexity: O(S), we need additional space to store the trie.
*/
func longestCommonPrefixV2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	type trieNode struct {
		children  map[rune]*trieNode
		isWordEnd bool
	}

	var root *trieNode
	newTrieNode := func() *trieNode {
		n := &trieNode{}
		n.children = make(map[rune]*trieNode)
		n.isWordEnd = false

		return n
	}

	insert := func(s string) {
		if len(s) == 0 {
			s = " "
		}

		cur := root
		for _, c := range s {
			n := cur.children[c]
			if n == nil {
				n = newTrieNode()
				cur.children[c] = n
			}
			cur = n
		}

		cur.isWordEnd = true
	}

	root = newTrieNode()
	for _, s := range strs {
		insert(s)
	}

	prefix := ""
	cur := root
	if len(cur.children) != 1 {
		return prefix
	}

	var r rune
	for (len(cur.children) == 1) && !cur.isWordEnd {
		for r = range cur.children {
			prefix += string(r)
		}

		cur = cur.children[r]
	}

	if prefix == " " {
		prefix = ""
	}

	return prefix
}
