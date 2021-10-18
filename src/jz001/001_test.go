package jz001

import "testing"

func Solution(input [][]int, target int) bool {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if target == input[i][j] {
				return true
			}
		}
	}
	return false
}

func Solution2(input [][]int, target int) bool {
	// 选择左下角元素, i 表示纵向index，j表示横向index
	i := len(input)
	j := 0
	for i > 0 && j < len(input[0]) {
		ele := input[i][j]
		if ele < target {
			i--
		} else if ele > target {
			j++
		} else {
			return true
		}
	}
	return false
}

func TestSolution(t *testing.T) {
}

