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
func reverseString(str string) string {

    var sb strings.Builder
    for i := len(str) - 1; i >= 0; i-- {
        if str[i] != ' ' {
            sb.WriteByte(str[i])
        }
    }
    return sb.String()
}


func main() {
    str1 := "A man a plan a canal Panama"
    // str2 := "not a not a not"
	reversed := reverseString(str1)
    fmt.Println("The reverse version of " + str1 + " is: " + reversed)

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
