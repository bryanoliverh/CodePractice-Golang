package main

import (
    "fmt"
    "strings"
)

func isPalindromeStr(str string) bool {
    str = strings.ToLower(strings.ReplaceAll(str, " ", ""))

    for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
        if str[i] != str[j] {
            return false
        }
    }
    return true
}

func main() {
    str1 := "A man a plan a canal Panama"
    // str2 := "not a not a not"

    if isPalindromeStr(str1) {
        fmt.Println(str1, "is a palindrome.")
    } else {
        fmt.Println(str1, "is not a palindrome.")
    }

    // if isPalindromeStr(str2) {
    //     fmt.Println(str2, "is a palindrome.")
    // } else {
    //     fmt.Println(str2, "is not a palindrome.")
    // }
}