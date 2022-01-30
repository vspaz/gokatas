package collections

func getBiggestNumInArray(nums []int) int {
	if len(nums) == 0 {
		panic("array size can't be 0")
	}
	if len(nums) == 1 {
		return nums[0]
	}
	biggestNum := nums[0]
	for _, num := range nums {
		if num > biggestNum {
			biggestNum = num
		}
	}
	return biggestNum
}
