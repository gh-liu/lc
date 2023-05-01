package main

func main() {
}

// 15
// 0
// [-1,0,0,1,1,2,2,3,3,4,4,5,5,6,6]
// [1,1,1,1,1,1,1,0,0,0,0,0,0,0,0]
func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	t := make(map[int][]int)
	for e, m := range manager {
		t[m] = append(t[m], e)
	}

	var dfs func(int) int
	dfs = func(i int) int {
		var ret int
		for _, sibling := range t[i] {
			i2 := dfs(sibling)
			if i2 > ret {
				ret = i2
			}
		}
		return informTime[i] + ret
	}

	return dfs(headID)
}
