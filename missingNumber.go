func missingNumber(nums []int) int {
	a := len(nums)

	totalSum := (a * (a+1))/2
	realSum := 0
	for _, num in range nums{
	realSum+= num
	}

	return totalSum - realSum

}
