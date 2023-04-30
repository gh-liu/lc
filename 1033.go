package main

func main() {
}

func numMovesStones(a int, b int, c int) []int {
	// order
	x := min(min(a, b), c)
	z := max(max(a, b), c)
	y := a + b + c - x - z

	res := []int{2, z - x - 2}
	if (z-y) == 1 && (y-x) == 1 {
		res[0] = 0
	} else if (z-y) <= 2 || (y-x) <= 2 {
		res[0] = 1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
