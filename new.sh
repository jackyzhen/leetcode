#!/usr/bin/env bash

# Generate template go for leetcode problems
# Usage: ./new.sh https://leetcode.com/problems/count-and-say/description/

PROBLEM_URL=$1

PROBLEM=$(echo $1 | sed -r 's/.*problems\/([a-zA-Z\-]+).*/\1/')

mkdir $PROBLEM


TEMPLATE=$(cat <<-END
// $PROBLEM_URL

package main

import "fmt"

func funcUnderTest(test []int, target int) int {
    return 0
}

func main() {
    tt := []struct {
        test     []int
        target   int
        expected int
    }{
        {
            []int{1, 3, 5, 6}, 5, 2,
        },
    }

    for _, t := range tt {
        actual := funcUnderTest(t.test, t.target)
        if t.expected != actual {
            panic(fmt.Sprintf("test: %v, target: %v, expected %v, actual %v", t.test, t.target, t.expected, actual))
        }
    }
}
END
)

echo "$TEMPLATE" > "$PROBLEM/main.go"
vim "$PROBLEM/main.go"
