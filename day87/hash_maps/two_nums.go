package hash_maps

func RunTwoNums(in []int, target int) []int {

	numbers := make(map[int]int)

	for k, v := range in {
		complement := target - v
		if _, ok := numbers[complement]; ok {
			return []int{k, numbers[complement]}
		}
		numbers[v] = k
	}

	return []int{-1, -1}
}
