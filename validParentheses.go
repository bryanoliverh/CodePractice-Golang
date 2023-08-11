package main

import "fmt"

func isValid(s string) bool {
    stack := []rune{} // Create an empty stack to store opening brackets
    mapping := map[rune]rune{ // Create a mapping of closing brackets to their corresponding opening brackets
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for _, char := range s { // Iterate over each character in the input string
        if char == '(' || char == '{' || char == '[' { // If the character is an opening bracket
            stack = append(stack, char) // Push it onto the stack
        } else if char == ')' || char == '}' || char == ']' { // If the character is a closing bracket
            if len(stack) == 0 || stack[len(stack)-1] != mapping[char] {
                // If the stack is empty or the top of the stack does not match the corresponding opening bracket
                return false // The input string is not valid
            }
            stack = stack[:len(stack)-1] // Pop the top element from the stack
        }
    }

    return len(stack) == 0 // If the stack is empty, the input string is valid
}

func main() {
    input := "({[]})"
    fmt.Println(isValid(input)) // Output: true
}
