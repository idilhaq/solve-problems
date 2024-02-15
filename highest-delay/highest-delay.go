package highestdelay

func HighestDelay(arr [][]int) string {
	var currDelay, currKey int
	for i, valArr := range arr {
		if i == 0 {
			currKey = valArr[0]
			currDelay = valArr[1]
			continue
		}

		diff := valArr[1] - currDelay
		if diff > currDelay {
			currKey = valArr[0]
			currDelay = diff
		}
	}

	alphabet := "abcdefghijklmnopqrstuvwxyz"

	return alphabet[currKey-1 : currKey]
}
