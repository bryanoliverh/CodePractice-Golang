func sortArray(nums []int) []int {
    quicksort(nums, 0, len(nums)-1)
    return nums
}

func quicksort(nums []int, low, high int) {
    if low < high {
        p := partition(nums, low, high)
        quicksort(nums, low, p-1)
        quicksort(nums, p+1, high)
    }
}

func partition(nums []int, low, high int) int {
    initial := nums[high]
    i := low
    for j := low; j < high; j++ {
        if nums[j] < initial {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
    }
    nums[i], nums[high] = nums[high], nums[i]
    return i
}
