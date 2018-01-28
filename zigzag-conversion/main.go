// https://leetcode.com/problems/zigzag-conversion/description/
package main

import (
	"bytes"
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	rows := make([][]string, numRows)
	rc := 0
	isDown := false
	for _, c := range s {
		rows[rc] = append(rows[rc], string(c))
		if rc == numRows-1 || rc == 0 {
			isDown = !isDown
		}
		if isDown {
			rc += 1
		} else {
			rc -= 1
		}
	}
	var buf bytes.Buffer
	for i := 0; i < numRows; i++ {
		for _, c := range rows[i] {
			buf.WriteString(c)
		}
	}
	return string(buf.Bytes())
}

func main() {
	fmt.Println(convert("PAHNAPLSIIGYIR", 3))
}
