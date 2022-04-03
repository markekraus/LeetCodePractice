package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	perm := []byte(s1)
	target := []byte(s2)
	permtable := make([]int, 26, 26)
	for _, v := range perm {
		permtable[v-97]++
	}
	for i := 0; i < len(s2); i++ {
		visited := make([]int, 26, 26)
		run := 0
		for run = 0; run < len(s1); run++ {
			if i+run >= len(s2) {
				return false
			}
			visited[target[i+run]-97]++
			if permtable[target[i+run]-97] < 1 || visited[target[i+run]-97] > permtable[target[i+run]-97] {
				break
			}
		}
		if run == len(s1) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Printf("%v\n", checkInclusion("ab", "eidbaooo"))
	fmt.Printf("%v\n", checkInclusion("ab", "eidboaoo"))
}
