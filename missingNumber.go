func missingNumber(nums []int) int {
    n := len(nums)
    totalSum := n * (n + 1) / 2

    for _, num := range nums {
        totalSum -= num
    }

    return totalSum
}
