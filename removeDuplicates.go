func removeDuplicates(nums []int) int {
    index := 1
    // sort.Sort(sort.Reverse(sort.IntSlice(nums)))
    sort.Sort(sort.IntSlice(nums))
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i - 1] {
            nums[index] = nums[i]
            index++
        }
    }
    return index
}
