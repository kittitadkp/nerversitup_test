package integer

func FindOddOccurrence(nums []int) int {
	countMap := make(map[int]int)

	for _, num := range nums {
		countMap[num]++
	}

	for num, count := range countMap {
		if count%2 != 0 {
			return num
		}
	}

	return -1
}
