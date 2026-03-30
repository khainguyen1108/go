package array

import "sort"

func UnmarkedSumArray(nums []int, queries [][]int) []int64 {
	arrMarked := make([]int, len(nums))
	result := make([]int64, len(queries))
	for i := 0; i < len(queries); i++ {
		arrMarked[queries[i][0]] = 1
		for t := 0; t < queries[i][1]; t++ {
			idx := findMin(nums, arrMarked)
			arrMarked[idx] = 1
		}
		result[i] = calcSumArr(nums, arrMarked)
	}

	return result
}

func findMin(nums []int, arrMarked []int) int {
	min := 0
	indexMin := 0

	for idm, v := range nums {
		if arrMarked[idm] != 1 {
			min = v
			indexMin = idm
			break
		}
	}
	for idm, v := range nums {
		if v < min && arrMarked[idm] != 1 {
			min = v
			indexMin = idm
		}
	}
	return indexMin
}

func calcSumArr(nums []int, arrMarked []int) int64 {
	sum := 0
	for idm, v := range nums {
		if arrMarked[idm] == 0 {
			sum += v
		}
	}
	return int64(sum)
}

func unmarkedSumArray(nums []int, queries [][]int) []int64 {
	n := len(nums)
	marked := make([]bool, n)

	// Tính tổng ban đầu
	totalSum := int64(0)
	for _, v := range nums {
		totalSum += int64(v)
	}

	// Sắp xếp indices theo value tăng dần (tie-break: index nhỏ hơn trước)
	sortedIdx := make([]int, n)
	for i := range sortedIdx {
		sortedIdx[i] = i
	}
	sort.Slice(sortedIdx, func(a, b int) bool {
		if nums[sortedIdx[a]] == nums[sortedIdx[b]] {
			return sortedIdx[a] < sortedIdx[b]
		}
		return nums[sortedIdx[a]] < nums[sortedIdx[b]]
	})

	result := make([]int64, len(queries))
	ptr := 0 // con trỏ vào sortedIdx, chỉ đi tiến, không lùi

	for i, q := range queries {
		markIdx, k := q[0], q[1]

		// Đánh dấu index theo query
		if !marked[markIdx] {
			marked[markIdx] = true
			totalSum -= int64(nums[markIdx])
		}

		// Tìm k phần tử nhỏ nhất chưa marked
		// ptr không reset về 0 vì các phần tử trước ptr đã marked hết
		for k > 0 && ptr < n {
			idx := sortedIdx[ptr]
			if marked[idx] {
				ptr++
				continue
			}
			marked[idx] = true
			totalSum -= int64(nums[idx])
			ptr++
			k--
		}

		result[i] = totalSum
	}

	return result
}
