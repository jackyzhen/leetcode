#!/usr/bin/env bash

# Generate template js for leetcode problems
# Usage: ./new-js.sh https://leetcode.com/problems/count-and-say/description/

PROBLEM_URL=$1

PROBLEM=$(echo $1 | sed -r 's/.*problems\/([a-zA-Z\-]+).*/\1/')

mkdir $PROBLEM

vim "$PROBLEM/index.js"
