func searchRange(nums []int, target int) []int {
    // Find the leftmost occurrence of the target
    left := binarySearch(nums, target, true)
    
    // Find the rightmost occurrence of the target
    right := binarySearch(nums, target, false)

    if left <= right {
        // If the target is found, return the range
        return []int{left, right}
    } else {
        // If the target is not found, return [-1, -1]
        return []int{-1, -1}
    }
}

func binarySearch(nums []int, target int, isFirst bool) int {
    low := 0
    high := len(nums) - 1
    result := -1

    for low <= high {
        mid := low + (high - low) / 2

        if nums[mid] == target {
            // Update the result with the current mid index
            result = mid

            if isFirst {
                // If finding the leftmost occurrence, narrow the search range to the left
                high = mid - 1
            } else {
                // If finding the rightmost occurrence, narrow the search range to the right
                low = mid + 1
            }
        } else if nums[mid] < target {
            // If the current element is less than the target, search in the right half
            low = mid + 1
        } else {
            // If the current element is greater than the target, search in the left half
            high = mid - 1
        }
    }

    return result
}
