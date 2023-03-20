package main

import "fmt"
import "strconv"

func maxSubArray(nums []int) int {
    maxSum, currSum := nums[0], nums[0]
    startIndex, endIndex := 0, 0
    for i := 1; i < len(nums); i++ {
        if nums[i] > currSum+nums[i] {
            currSum = nums[i]
            startIndex = i
        } else {
            currSum = currSum + nums[i]
        }
        if currSum > maxSum {
            maxSum = currSum
            endIndex = i
        }
    }
    fmt.Printf("Indices %d to %d: ", startIndex, endIndex)
    for i := startIndex; i <= endIndex; i++ {
        fmt.Printf("%d ", nums[i])
    }
    fmt.Println()
    return maxSum
}

func main() {
    arr := []int{-7, 2, -3, 1, 2, 3, 6, -5, 10}
    fmt.Println("Largest Sum Subarray: " + strconv.Itoa(maxSubArray(arr)))
}