func searchRange(nums []int, target int) []int {
	output := []int{}
	finoutput := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			output = append(output, i)
		}
	}

	lengthOutput := len(output)
	if lengthOutput == 1 {
		finoutput[0] = output[0]
		finoutput[1] = output[0]
	} else if lengthOutput > 1 {
		finoutput[0] = output[0]
		finoutput[1] = output[lengthOutput-1]
	} else {
		finoutput[0] = -1
		finoutput[1] = -1
	}

	return finoutput
}
