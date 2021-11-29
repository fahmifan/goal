package stats

func Run(nums []int) int {
	// make pair of pairs start from pair of 2 to pair of n-1
	mapPairs := make(map[int][][]int)
	// minimum reversed pair is 3
	for i := 3; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			var pairs []int
			for k := j; k < j+i-1; k++ {
				if k >= len(nums) {
					break
				}
				pairs = append(pairs, nums[k])
			}
			if len(pairs) < i-1 {
				continue
			}
			pairKey := i - 1
			mapPairs[pairKey] = append(mapPairs[pairKey], pairs)
		}
	}

	// check each pairs, if there is a reversed pair
	// use, the first reversed pair found
	var reversedPair []int
	for _, pairs := range mapPairs {
		for i := 0; i < len(pairs); i++ {
			for j := i + 1; j < len(pairs); j++ {
				ok := checkReverse(pairs[i], pairs[j])
				if ok {
					reversedPair = pairs[i]
					break
				}
			}
		}
	}

	// find the max
	return max(reversedPair)
}

func max(nums []int) int {
	if nums == nil {
		return 0
	}
	res := nums[0]
	for _, num := range nums {
		if num > res {
			res = num
		}
	}

	return res
}

func checkReverse(x, y []int) bool {
	size := len(x)
	for i := 0; i < size; i++ {
		if x[i] != y[size-i-1] {
			return false
		}
	}

	return true
}
