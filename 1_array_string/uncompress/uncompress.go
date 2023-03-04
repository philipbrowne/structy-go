package main

import (
	"log"
	"strconv"
	"strings"
)

func main () {}

// Two Pointer Approach - O(n*m) Runtime, where n is the number of groups and m is the max number found in any group
func uncompress(s string) string {
	res := []string{}
	i := 0
	j := 0
	numStr := "0123456789"
	for j < len(s){
		if strings.Contains(numStr, string(s[j])) {
			j++
		} else {
			num, err := strconv.Atoi(string(s[i:j]))
			if err != nil {
				log.Fatal(err)
			}
			for n := 0; n < num; n++ {
				res = append(res, string(s[j]))
			}
			j++
			i = j
		}
	}
	return strings.Join(res, "")
}