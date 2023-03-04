package main

import (
	"strconv"
	"strings"
)

func main () {}

// Two Pointer Approach - O(n) Runtime
func compress (s string) string {
	res := []string{}
	i := 0
	j := 0
	s += "!"
	for j < len(s) {
		if string(s[i]) == string(s[j]) {
			j ++
		} else {
			count := j - i
			if count == 1 {
				res = append(res, string(s[i]))
			} else {
				res = append(res, strconv.Itoa(count), string(s[i]))
			}
			i = j
			j++
		}
	}
	return strings.Join(res, "")
}