package booyermoremajorityvote

// Boyer-Moore Majority Vote
func MajorityElements(nums []int) []int {
	count := make(map[int]int)
	result := make([]int, 0)

	for _, num := range nums {
		count[num]++
	}

	for num, freq := range count {
		if freq > len(nums)/3 {
			result = append(result, num)
		}
	}

	return result
}