package main

import "github.com/luo/leetcode-kit/leetcode"

// print sth to help you debug

func main() {
	root := leetcode.ArrayToTree([]int{5,4,8,11,-1,13,4,7,2,-1,-1,5,1})
	leetcode.PrintTree(root, "", true)
}
