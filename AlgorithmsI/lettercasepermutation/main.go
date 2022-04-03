package main

import "fmt"

func letterCasePermutation(S string) []string {
	ans := make([]string, 0, len(S))
	dfs([]byte(S), 0, &ans)
	return ans
}

func dfs(s []byte, i int, ans *[]string) {
	if i == len(s) {
		*ans = append(*ans, string(s))
		return
	}
	dfs(s, i+1, ans)
	if isAlphabet(s[i]) {
		s[i] ^= 32
		dfs(s, i+1, ans)
	}
}

func isAlphabet(c byte) bool {
	if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
		return true
	}
	return false
}

func main() {
	fmt.Printf("%v\n", letterCasePermutation("a123b4"))
}
