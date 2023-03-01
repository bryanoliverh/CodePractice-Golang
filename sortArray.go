func sortArray(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }

    initial := nums[0]
    var left, right []int

    for _, num := range nums[1:] {
        if num <= initial {
            left = append(left, num)
        } else {
            right = append(right, num)
        }
    }

    left = sortArray(left)
    right = sortArray(right)

    return append(append(left, initial), right...)
}
