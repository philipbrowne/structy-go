package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main () {
	compressed := compress(os.Args[1])
	fmt.Println(compressed)
}

// compress("ssssbbz") -> "4s2bz"
// compress("ppoppppp") -> "2po5p"
// compress("ccaaatsss") -> "2c3a53s"
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