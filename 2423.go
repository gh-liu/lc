package main

func main() {
	equalFrequency("abcc")
	equalFrequency("aacc")
}

func equalFrequency(word string) bool {
	var idx [26]int
	for _, v := range word {
		i := int8(v) - 97
		idx[i]++
	}
	for i := 0; i < 26; i++ {
		if idx[i] == 0 {
			continue
		}
		idx[i]--
		frequency := make(map[int]bool)
		for _, f := range idx {
			if f > 0 {
				frequency[f] = true
			}
		}
		if len(frequency) == 1 {
			return true
		}
		idx[i]++
	}

	return false
}
