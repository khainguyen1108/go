package array

func xorAfterQueries(nums []int, queries [][]int) int {
	MOD := 1_000_000_007

	for i := 0; i < len(queries); i++ {
		t := queries[i][0]
		for t <= queries[i][1] {
			nums[t] = (queries[i][3] * nums[t]) % MOD
			t += queries[i][2]
		}
	}
	res := 0

	for i := 0; i < len(nums); i++ {
			res ^= nums[i]
	}
	return res
}
