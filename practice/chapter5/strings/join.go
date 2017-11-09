package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []string{"oo", "00", "11"}
	fmt.Println(join("->", "aa", "bb", "cc", "dd"))
	fmt.Println(join("->", s...))
}

func join(sep string, vals ...string) string {
	var buf bytes.Buffer
	for index, val := range vals {
		buf.WriteString(val)
		if index != (len(vals)-1) {
			buf.WriteString(sep)
		}
	}
	return buf.String()
}
