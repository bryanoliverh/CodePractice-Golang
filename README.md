# List of some of my LeetCode practice codes!

***1. [isPalindrome.go](https://github.com/bryanoliverh/LeetCode-Practice-Golang/blob/main/sortArray.go)***

The function isPalindrome takes an integer x as input and it will return a boolean which indicates whether x is a palindrome or not. A palindrome is a number that reads the same backward as forward.
so for example:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

The function first checks if x is negative or ends with a zero, except when x is zero itself. If either condition is true, the function immediately returns false because such numbers cannot be palindromes.

Next, the function initializes a variable rev to zero. It then enters a loop where it repeatedly extracts the last digit of x and adds it to rev after shifting the digits of rev one position to the left. This effectively reverses the digits of x and stores the result in rev. At the same time, the function removes the last digit of x by integer division with 10.

The loop continues as long as x is greater than rev, since the reversal of the first half of x is not complete until x becomes smaller than rev. If x and rev have an odd number of digits, the middle digit is ignored because it is the same in both numbers.

After the loop completes, the function checks whether x is equal to rev or to rev/10. The second case arises when x has an odd number of digits, and the middle digit has already been removed from rev.

If either of these conditions is true, the function returns true, indicating that x is a palindrome. Otherwise, it returns false.

***2. [sortArray.go](https://github.com/bryanoliverh/LeetCode-Practice-Golang/blob/main/sortArray.go)***

Sort Array Function in GO without built-in functions!


The function takes 1 list input and returns the sorted list without any built in functions.

Time Complexity: O(n log n)
Space Complexity: O(n)

The function implements quick sort algorithm which takes slice of integers with the variable name of nums, and the range of indices low and high which will be sorted. If low is less than high, the slice will be partitioned using the partition function, and recursively calls itself on the left and right sub-arrays.

Another function is the partition function which is used to do the partition of the input slice of integers nums. It will choose the last element in the slice as the initial element that will be use as the pivot, and iterates over the slice from low to high-1. If an element is less than the initial element, it swaps it with the current index i, and increments i. Finally, it swaps the pivot element with the element at index i and returns i.

