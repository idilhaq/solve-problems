package maxpair

import (
	"fmt"
)

func MaxPair() {
	input := [][]int{{1, 2}, {2, 3}, {2, 1}, {2, 1}, {1, 2}, {2, 1}, {1, 2}, {2, 2}, {2, 2}, {3, 2}}

	pairSet := make(map[[2]int]bool)
	pairSetDetail := make(map[[2]int]int)

	maxPairs := 0

	for _, pair := range input {
		pairStr := [2]int{pair[0], pair[1]}
		pairReverse := [2]int{pair[1], pair[0]}

		if !pairSet[pairStr] {
			pairSet[pairStr] = true
		}

		if pairSet[pairReverse] {
			if pairSetDetail[pairStr] == 1 {
				pairSetDetail[pairStr] = 0
				continue
			}

			maxPairs++
			pairSetDetail[pairStr]++
		}
	}

	fmt.Println("Expected max pair is:", maxPairs)
}
