package hashtable

func GetDistance1(arr []int) []int64 {
	arrRes := make([]int64, len(arr))
	for i := 0; i < len(arr); i++ {
		weight := 0
		for j := 0; j < len(arr); j++ {
			if arr[j] == arr[i] && i != j {
				weight += abs1(j - i)
			}
		}
		arrRes[i] = int64(weight)
		weight = 0
	}
	return arrRes
}

func abs1(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GetDistance(arr []int) []int64 {
	n := len(arr)
	res := make([]int64, n)

	// map value -> list of indices
	mp := map[int][]int{}

	for i, v := range arr {
		mp[v] = append(mp[v], i)
	}

	for _, pos := range mp {

		m := len(pos)

		prefix := make([]int64, m+1)

		for i := 0; i < m; i++ {
			prefix[i+1] = prefix[i] + int64(pos[i])
		}

		for i := 0; i < m; i++ {

			p := int64(pos[i])

			left := p*int64(i) - prefix[i]
			right := (prefix[m] - prefix[i+1]) - p*int64(m-i-1)

			res[pos[i]] = left + right
		}
	}

	return res
}
