package main

// func main() {
// }

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	step := func(left int, right int) int {
		if fruits[right][0] <= startPos {
			return startPos - fruits[left][0]
		} else if fruits[left][0] >= startPos {
			return fruits[right][0] - startPos
		} else {
			return min(abs(startPos-fruits[right][0]), abs(startPos-fruits[left][0])) + fruits[right][0] - fruits[left][0]
		}
	}

	sum := 0
	ans := 0

	left := 0
	right := 0

	// 右指针开始移动
	for right < len(fruits) {
		sum += fruits[right][1]
		// 如果左指针在右指针左侧，且从开始位置到两者之间的步数大于k，左指针右进
		for left <= right && step(left, right) > k {
			sum -= fruits[left][1]
			left++
		}

		// 记录最大值
		ans = max(ans, sum)
		right++
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
