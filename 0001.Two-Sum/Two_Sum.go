package leetcode_go

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		other := target - num
		if _, ok := m[other]; ok {
			return []int{m[other], i}
		}
		m[num] = i
	}
	return nil
}
