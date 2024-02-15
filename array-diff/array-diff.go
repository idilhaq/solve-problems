package arraydiff

func ArrayMultiplier(arr1, arr2 []int, counter int) int {
	var diff int

	if counter > len(arr1) {
		return diff
	}

	for i := range arr1 {
		diff += arr1[i] - arr2[i]
	}

	counter++

	return diff + ArrayMultiplier(append(arr1[1:], arr1[:1]...), arr2, counter)
}
