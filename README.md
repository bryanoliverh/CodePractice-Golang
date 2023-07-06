# List of some of my results of my weekly exercise from LeetCode, HackerRank, and other challenges!

Please do feel free to ask questions or any inquiries about the code or the logic behind the code through creating issues in this repo or by email. The short summaries of these codes are written below:

1. [compress.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/compress.go)

 - Time Complexity: O(n)
   
   Space Complexity: O(1)

 - **Input and Output Examples:**
 
   **Example 1:**

   Input: chars = ["a","a","b","b","c","c","c"] 
   Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"] 
   Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3". 
   
   **Example 2:** 

   Input: chars = ["a"] 
   Output: Return 1, and the first character of the input array should be: ["a"] 
   Explanation: The only group is "a", which remains uncompressed since it's a single character. 
  
   **Example 3:**

   Input: chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"] 
   Output: Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"]. 
   Explanation: The groups are "a" and "bbbbbbbbbbbb". This compresses to "ab12". 

 - It will first check if the length of chars is 0 or 1 and it will return 0 or 1 respectively, since there is no need to compress a sequence of 0 or 1 character.

 - The function initializes comp variable to be the first character in chars, count to be 1, and compressIdx to be 0 and it will then enter a for loop that iterates over the characters in chars, starting from the second character. If the current character is the same as the previous one (comp), the function increments the count variable by 1.

 - If the current character is different from the previous one, the function writes the previous character (comp) to the compressIdx position in chars and increments compressIdx by 1.If the count variable is greater than 1, it means there are repeated characters, so the function converts the count to a string using strconv.Itoa(count) and writes each digit of the count to the compressIdx position in chars, and increments compressIdx by the length of the count string. The function then updates comp to be the current character and resets count to 1. After the loop is finished, the function returns the value of compressIdx, which represents the length of the compressed sequence.

2. [isPalindrome.go](https://github.com/bryanoliverh/LeetCode-Practice-Golang/blob/main/sortArray.go)

  - Time Complexity: O(log(n))
  
    Space Complexity: O(n)

  - The function isPalindrome takes an integer x as input and it will return a boolean which indicates whether x is a palindrome or not. A palindrome is a number that reads the same backward as forward.
so for example:

  - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: x = 121
    
    Output: true
    
    Explanation: 121 reads as 121 from left to right and from right to left.

  - The function first checks if x is negative or ends with a zero, except when x is zero itself. If either condition is true, the function immediately returns false because such numbers cannot be palindromes.

  - Afterward, the function initializes a variable rev to zero and it will enter a loop where it repeatedly extracts the last digit of x and adds it to rev after shifting the digits of rev one position to the left. Therefore, this effectively reverses the digits of x and stores the result in rev. At the same time, the function removes the last digit of x by integer division with 10.

  - The loop will continue as long as x is greater than rev, since the reversal of the first half of x is not complete until x becomes smaller than rev. If x and rev have an odd number of digits, the middle digit is ignored because it is the same in both numbers. After the loop completes, the function checks whether x is equal to rev or to rev/10 and the second case arises when x has an odd number of digits, and the middle digit has already been removed from rev.
  - The expression x == reverse || x == reverse / 10 is checking whether the input integer x is a palindrome or not while at the end of the for loop, the variable reverse will contain the reverse of the first half of x. For example, if x is 1234321, then reverse will be 1234. Then it will compare the first half of x with the second half to check if it's a palindrome. If x has an odd number of digits, we can ignore the middle digit, which will be equal to itself. If x has an even number of digits, we need to compare the first half of x with the second half minus the last digit. For example, if x is 1221, then reverse will be 12, and we need to ignore the last digit in reverse. Therefore, the expression x == reverse / 10 checks if x is a palindrome when it has an even number of digits.
  - If either of these conditions is true, the function returns true, indicating that x is a palindrome. Otherwise, it returns false.

  
3. [sortArray.go](https://github.com/bryanoliverh/LeetCode-Practice-Golang/blob/main/sortArray.go)
  - Time Complexity: O(n log n)
    
    Space Complexity: O(n)

  - Sort Array Function in GO without built-in functions!

  - The function takes 1 list input and returns the sorted list without any built in functions.


  - The function implements quick sort algorithm which takes slice of integers with the variable name of nums, and the range of indices low and high which will be sorted. If low is less than high, the slice will be partitioned using the partition function, and recursively calls itself on the left and right sub-arrays.

 - Another function is the partition function which is used to do the partition of the input slice of integers nums. It will choose the last element in the slice as the initial element that will be use as the pivot, and iterates over the slice from low to high-1. If an element is less than the initial element, it swaps it with the current index i, and increments i. Finally, it swaps the pivot element with the element at index i and returns i.

4. [twoSum.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/twoSum.go)

  - Time Complexity: O(n)
   
    Space Complexity: O(n)
  - Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.  You may assume that each input would have exactly one solution, and you may not use the same element twice. 
  - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: nums = [2,7,11,15], target = 9 
    
    Output: [0,1] 
    
    Explanation: Because nums[0] + nums[1] == 9, we return [0, 1]. 
  - The twoSum function takes an integer slice array and a target integer target as its input, and returns a slice of two integers that represent the indices of the two numbers in the input array that add up to the target value. If there are no such indices, an empty slice is returned. 
  - The function uses a hash map called seenNums to keep track of the numbers it has already seen while iterating through the array. In each iteration, it calculates the difference between the target value and the current number in the array and checks whether the difference is already in the hash map. If it is, the function returns a slice with the current index i and the index of the difference j that was previously stored in the hash map. If the difference is not in the hash map, the function stores the current number and its index in the hash map. 
  - seenNums map stores the numbers in the input array as keys and their corresponding indices as values. The syntax j, ok := seenNums[potentialMatch] will simultaneously checks whether the value exists in the map and retrieves its value. The value of found will be true if potentialMatch is a key in the seenNums map, and false otherwise. The value of j will be the value associated with the key potentialMatch if it exists in the map, and the zero value for its type otherwise. 
  
 5. [intToRoman_romanToInt.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/intToRoman_romanToInt.go)
   -  Time Complexity: O(n)
   
      Space Complexity: O(1)
  - Given an integer, convert it to a roman numeral. 
  - **Input and Output Examples:**
  
     **Example 1:** 
     
     Input: num = 3 
     
     Output: "III" 
     
     Explanation: 3 is represented as 3 ones. 

     **Example 2:** 
     
     Input: num = 58 
     
     Output: "LVIII" 
     
     Explanation: L = 50, V = 5, III = 3. 
   - romanToInt is the reverse function that will take roman characters and map it to Integer/common numbers.
   
   
6. [threeSumClosest.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/threeSumClosest.go)
  -  Time Complexity: O(n^2)
   
      Space Complexity: O(1)
      
  - This function finds the closest sum of three numbers in the given input slice nums to a target number target. 
  - The function starts by sorting the input slice. Then, it iterates over the first len(nums) - 2 elements of the sorted slice nums, while keeping track of the closest sum seen so far and the distance of the closest sum to the target. 

  - For each element nums[i] in the outer loop, the function initializes two pointers left and right that point to the next and last elements, respectively. It then iterates over the slice from the left and right pointers until they meet. 

 
 

  - At each iteration, the function calculates the sum of the three elements at indices i, left, and right. If this sum is less than the target, the left pointer is incremented, otherwise, the right pointer is decremented. If the sum equals the target, the function returns the target. 

  - For each calculated sum, the function checks if it is closer to the target than the closest sum seen so far. If so, it updates the closest sum and the distance accordingly. 

  - Finally, the function returns the closest sum found. 
  
7. [mergeTwoList.go](  https://github.com/bryanoliverh/CodePractice-Golang/blob/main/mergeTwoList.go)
  -  Time Complexity: O(m+n), where m and n are the lengths of list1 and list2 
   
      Space Complexity: O(1)
      
  - **Input and Output Examples:**
  
    **Example 1:** 
    
    Input: list1 = [1,2,4], list2 = [1,3,4]
    
    Output: [1,1,2,3,4,4]

    **Example 2:**

    Input: list1 = [], list2 = []
    
    Output: []
  
  
  - The mergeTwoLists function takes in two singly-linked lists list1 and list2 and merges them into a new singly-linked list. To create the new list, a dummy node is created as a placeholder, and a pointer tail is set to it.

  - Then, a loop is initiated to compare each node in both lists. If the value of the current node in list1 is less than or equal to the value of the current node in list2, the tail node's Next pointer is set to point to the current node in list1, and list1 moves on to the next node in the list. Otherwise, the tail node's Next pointer is set to point to the current node in list2, and list2 moves on to the next node in the list. The tail node is then moved to the next node in the new list.

  - After the loop has completed, any remaining nodes in either list1 or list2 are added to the new list. If there are still nodes remaining in list1, the tail node's Next pointer is set to point to list1. If there are still nodes remaining in list2, the tail node's Next pointer is set to point to list2.

  - Finally, the function returns the Next pointer of the dummy node. This is because the first node in the new list is the Next node of the dummy node, which was created only as a placeholder and has no other significance in the new list. By returning the Next pointer of the dummy node, we are effectively returning the first node of the new list.

8. [removeDuplicates.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/removeDuplicates.go)
  -  Time Complexity: O(n log n) which is dominated by the sorting function.
   
     Space Complexity: O(1)
      
  - **Input and Output Examples:**
  
    **Example 1:**

    Input: nums = [1,1,2]
    
    Output: 2, nums = [1,2,_]
    
    Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
    It does not matter what you leave beyond the returned k (hence they are underscores).
    
    
    **Example 2:**

    Input: nums = [0,0,1,1,1,2,2,3,3,4]
    
    Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
    
    Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
    
    It does not matter what you leave beyond the returned k (hence they are underscores).


9. [lengthOfLastWord.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/lengthOfLastWord.go)
  -  Time Complexity: O(n)
   
     Space Complexity: O(n)
      
  - **Input and Output Examples:**

    **Example 1:**
    
    Input: s = "Hello World"
    
    Output: 5
    
    Explanation: The last word is "World" with length 5.
    
    **Example 2:**

    Input: s = "   fly me   to   the moon  "
    
    Output: 4
    
    Explanation: The last word is "moon" with length 4.


  - Given a string s consisting of words and spaces, return the length of the last word in the string.
  
  - A word is a maximal substring consisting of non-space characters only.
  
  - The function loops through each character in the input string and it will check if the character is a space (ASCII code 32). If it is a space, it means that it is going to be a new word and the new string will be emptied. Otherwise, the character is added to the current finStr variable.
  
 10. [interval.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/interval.go)
  -  The task is to combine overlapping intervals in a given list to create a new list that contains only intervals that do not overlap with each other.

  -  Time Complexity: O(n log n)
   
     Space Complexity: O(n)
      
  - **Input and Output Examples:**

    **Example 1:**
    
    Input: {1, 10}, {2, 9}, {5, 11}, {3, 4}
    
    Output: (1, 11)
 11. [determineIdenticalBinaryTree.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/determineIdenticalBinaryTree.go)
  -  The problem asks to determine if two binary trees are structurally identical, meaning they have the same shape and nodes with the same values.
  
  -  Compare each node of the two binary trees, starting from their roots and going down to their leaves. If the nodes being compared are equal, the function recursively calls itself on the left and right children of the nodes. If at any point during the recursion, two nodes being compared are not equal, the function returns false. If the function completes the recursion without returning false, it means that the two trees are identical, and the function returns true.

  -  Time Complexity: O(n)
   
     Space Complexity: O(n)
      
  - **Input and Output Examples:**

    **Example 1:**
    
    Input: Tree 1:
            1
           │   2
           │  │   4
           │      5
               3
              │   6
           Tree 2:
            1
           │   2
           │  │   4
           │      5
               3
              │   6

    
    
    Output: The two binary trees are identical.
 
  12. [isPalindromeStr.go]( https://github.com/bryanoliverh/CodePractice-Golang/blob/main/isPalindromeStr.go)
  -  Check whether the given string is a palindrome or not.

  -  Time Complexity: O(n)
   
     Space Complexity: O(n)
      
  - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: "A man a plan a canal Panama is a palindrome."
    
    Output: not a not a not is not a palindrome.
    

    **Example 2:**
    
    Input: "not a not a not"
    
    Output: not a not a not is not a palindrome.
    
13. [maxSubArray.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/maxSubArray.go)

  -  Create a function to compute the largest sum subarray of the input array.

  -  Time Complexity: O(n)
   
     Space Complexity: O(1)
      
  - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: [-2, 1, -3, 4, -1, 2, 1, -5, 4]
    
    Output: Indices 3 to 6: 4 -1 2 1  Largest Sum Subarray: 6
            
    **Example 2:**
     
     Input: [1, 2, 3, 4, 5]

     Output: Indices 0 to 4: 1 2 3 4 5 Largest Sum Subarray: 15

14. [printPeople.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/printPeople.go)

  -  Create a function to practice using linkedlist, struct, list, and Map.

  -  Time Complexity: O(n)
   
     Space Complexity: O(n)
      
  - **Input and Output Examples:**
    
    **Example 1:**
        
    Output: 
    ```     
            Name of All Registered People

            ID: 1, Name: John, Age: 17, Location: New York
            
            ID: 2, Name: Jane, Age: 22, Location: Los Angeles
            
            ID: 3, Name: Bob, Age: 30, Location: Chicago
            
            Name of People Above 20 yo:Name: John, Age: 17
            
            Name of People Above 20 yo:Name: Jane, Age: 22
            
            Name of People Above 20 yo:Name: Bob, Age: 30
            
15. [longestCommonPrefix.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/longestCommonPrefix.go)

  -  Create a function to find the longest common prefix string amongst an array of strings.


  -  Time Complexity: O(n*m) (m= the length of the shortest string in the strs slice for example the length of the longest common prefix, n= the total number of characters in all the strings in the strs slice).
   
     Space Complexity: O(1)
      
  - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: strs = ["flower","flow","flight"]
    
    Output: "fl"
    
    **Example 2:**
    
    Input: strs = ["dog","racecar","car"]
    
    Output: ""
 
 16. [longestCommonSubstring.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/longestCommonSubstring.go)

  -  Create a function to find the longest common substring string contained amongst an array of strings.

  -  Time Complexity: O(m*n^2) (m= the length of the first string in the strs slice, n refers to the total number of characters in all the strings in the strs slice).
   
     Space Complexity: O(m) (m= the length of the first string in the strs slice)
      
  - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: strs = ["helloeworld", "oewor", "worhelloe"]
    
    Output: "wor"        
 
 17. [addBinary.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/addBinary.go)

  -  Create a function to calculate an input from two binary strings.

  -  Time Complexity: O(n) 
   
     Space Complexity: O(n) 

 - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: a = "11", b = "1"
    
    Output: "100"   
        
    **Example 2:**
    
    Input: a = "1010", b = "1011"
    
    Output: "10101"   

 18. [indexFirstOccurence.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/indexFirstOccurence.go)

  -  Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

  -  The loop condition i <= len(haystack)-len(needle) ensures that the loop will only iterate up to the point where the remaining characters in haystack are greater than or equal to the length of needle. This is because there's no point in checking substrings that are shorter than the needle string. Inside the loop, haystack[i:i+len(needle)] is used to extract a substring of haystack starting from index i and with a length equal to the length of needle. This substring is then compared with the needle string using the equality operator ==. If the substring matches the needle string, the function immediately returns the value of i, indicating the index of the first occurrence of needle in haystack. If no match is found within the loop, the function will continue to the next iteration until all possible substrings of haystack have been checked. If no match is found at all, the function will reach the end and return -1 to indicate that needle is not part of haystack.

  -  Time Complexity: O((N-M)M), where N is the length of haystack and M is the length of needle.
   
     Space Complexity: O(1) 

 - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: haystack = "sadbutsad", needle = "sad"
    
    Output: 0
    
    Explanation: "sad" occurs at index 0 and 6. The first occurrence is at index 0, so we return 0.  
        
    **Example 2:**
    
    Input: haystack = "leetcode", needle = "leeto"
    
    Output: -1
    
    Explanation: "leeto" did not occur in "leetcode", so we return -1.  
 
 19. [removeElement.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/removeElement.go)

  -  Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val. Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums. Return k.

  -  Time Complexity: O(n), where n is the length of nums.
   
     Space Complexity: O(1) 

 - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: nums = [3,2,2,3], val = 3
    
    Output: 2, nums = [2,2,_,_]
    
    Explanation: Your function should return k = 2, with the first two elements of nums being 2. It does not matter what you leave beyond the returned k (hence they are underscores).
        
    **Example 2:**
    
    Input: nums = [0,1,2,2,3,0,4,2], val = 2
    
    Output: 5, nums = [0,1,4,0,3,_,_,_]
    
    Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4. Note that the five elements can be returned in any order. It does not matter what you leave beyond the returned k (hence they are underscores).

20. [reverseInt.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/reverseInt.go)

  -  Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0. Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

  -  Time Complexity: O(log(x))
   
     Space Complexity: O(1) 

 - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: x = 123
    
    Output: 321
    
    **Example 2:**
    
    Input: x = -123
    
    Output: -321
    
    **Example 3:**
    
    Input: x = 120
    
    Output: 21
    
 21. [removeNthFromEnd.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/removeNthFromEnd.go)

  -  Given the head of a linked list, remove the nth node from the end of the list and return its head.

  -  By using &ListNode{Next: head}, you create a new ListNode instance with its Next field pointing to the head node. This way, dummy becomes a pointer to this new node, and its Next field connects to the original list starting from the head node.

  -  Time Complexity: O(N), N is the number of nodes in the linked list
   
     Space Complexity: O(1) 

 - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: head = [1,2,3,4,5], n = 2
    
    Output: [1,2,3,5]
    
    **Example 2:**
    
    Input: head = [1], n = 1
    
    Output: []
    
    **Example 3:**
    
    Input: head = [1,2], n = 1
    
    Output: [1]
    
    
22. [addTwoNumbers.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/addTwoNumbers.go)

  -  You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list. You may assume the two numbers do not contain any leading zero, except the number 0 itself.

  -  In the given code snippet, dummy is a pointer to a new instance of ListNode, which acts as a dummy head node for the resulting linked list. It is initially empty, meaning it doesn't contain any value. current is another pointer that initially points to the dummy node. It is used to keep track of the current node in the resulting linked list while iterating through the input lists l1 and l2. By using a dummy head node, we can easily handle the edge case of an empty resulting linked list without needing special checks. The current pointer allows us to append new nodes to the end of the resulting linked list by updating its Next field with a new node.

  -  Time Complexity: O(max(N, M)), where N and M are the lengths of the input linked lists l1 and l2 respectively
   
     Space Complexity: O(max(N, M))
     
 - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: l1 = [2,4,3], l2 = [5,6,4]
    
    Output: [7,0,8]
    
    Explanation: 342 + 465 = 807.
    
    **Example 2:**
    
    Input: l1 = [0], l2 = [0]
    
    Output: [0]
    
    **Example 3:**
    
    Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
    
    Output: [8,9,9,9,0,0,0,1] 


23. [maxArea.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/maxArea.go)

  -  You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]). Find two lines that together with the x-axis form a container, such that the container contains the most water. Return the maximum amount of water a container can store. Notice that you may not slant the container.

  -  We search the minimum of left and right because the container can only hold water at the height of the minimum, for example left = 6, right = 8, the area would be 6 * the x axis which is the difference between both indexes.

  - The min and max functions are utility functions used to find the minimum and maximum values between two integers, respectively.

  -  Time Complexity: O(N), where N is the length of the height
   
     Space Complexity: O(1)
     
 - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: height = [1,8,6,2,5,4,8,3,7]
    
    Output: 49
    
    Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

    **Example 2:**
    
    Input: height = [1,1]
    
    Output: 1
    
24. [addNextNodeWithValue.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/addNextNodeWithValue.go)

  -  Golang code that adds a new node with a value of 5 (as an example) as the next node of any node with a value of 3 (as an example) in a linked list
  
  -  Time Complexity: O(N), where n is the number of nodes in the linked list
   
     Space Complexity: O(1)
     
  
  - **Input and Output Examples:**
    
    **Example 1:**
    
    Input:
    Linked list: 1 -> 2 -> 3 -> 4 -> 5
    Target value: 3
    New value: 5

    Output:
    Linked list: 1 -> 2 -> 3 -> 5 -> 4 -> 5

    **Example 2:**
    
    Input:
    Linked list: 3 -> 3 -> 3 -> 3 -> 3
    Target value: 3
    New value: 5

    Output:
    
    Linked list: 3 -> 5 -> 3 -> 3 -> 3 -> 3

25. [swapPairs.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/swapPairs.go)

  -  Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)
  
  -  After each iteration, the carrier will move to the first, which means it is going to start for the next pair.
  
  -  Time Complexity: O(N), where n is the number of nodes in the linked list
   
     Space Complexity: O(1)
     
  
  - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: head = [1,2,3,4]
    
    Output: [2,1,4,3]

    **Example 2:**
    
    Input: head = []
    
    Output: []

    **Example 3:**
    
    Input: head = [1]
    
    Output: [1]
    
 26. [deleteDuplicatesWithSorting.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/deleteDuplicatesWithSorting.go)

  -  Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

  -  Using the merge sort algorithm, it will find the middle part of the linked list by advancing fast two nodes at a time and slow one node at a time. Furthermore, it will sort the first half and the second half. Moreover, it will then merge both of the halves together.
  
  -  Time Complexity: O(n log n) in the average and worst cases, where n is the number of nodes in the linked list. 
   
     Space Complexity: O(N), where n is the number of nodes in the linked list
     
  
  - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: head = [1,1,2]
    
    Output: [1,2]

    **Example 2:**
    
    Input: head = [1,1,2,3,3]
    
    Output: [1,2,3]

27. [removeElements.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/removeElements.go)

  -  Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.

  -  By using a separate current variable to traverse the linked list and updating current.Next to skip the nodes with the target value, we can effectively removing the undesired nodes while preserving the reference to the original head. This way, we can return the modified linked list correctly at the end.
  
  -  Time Complexity: O(n), where n is the number of nodes in the linked list
   
     Space Complexity: O(1)
 - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: head = [1,2,6,3,4,5,6], val = 6
    
    Output: [1,2,3,4,5]
    
    **Example 2:**
    
    Input: head = [], val = 1
    
    Output: []

 28. [reverseList.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/reverseList.go)

  -  Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.

  -  nextnode := current.Next: This line creates a new variable nextnode and assigns it the value of the next node after the current node. It's essentially capturing the reference to the next node before we modify the current.Next pointer. 
  -  current.Next = prev: This line reverses the Next pointer of the current node and points it to the prev node. This effectively changes the direction of the linked list, making the current node point to its previous node.
  
  -  prev = current: This line updates the prev node to be the current node. Since we're moving forward in the linked list, the current node will become the previous node in the next iteration.
  
  -  current = nextnode: This line updates the current node to be the nextnode, which was captured in the first line. By moving current to the next node, we continue iterating through the linked list.
  
  -  Time Complexity: O(n), where n is the number of nodes in the linked list
   
     Space Complexity: O(1)
     
 - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: head = [1,2,3,4,5]
    
    Output: [5,4,3,2,1]
    
    **Example 2:**
    
    Input: head = [1,2]
    
    Output: [2,1]
   
 29. [getIntersectionNode.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/getIntersectionNode.go)

  -  Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect. If the two linked lists have no intersection at all, return null. The test cases are generated such that there are no cycles anywhere in the entire linked structure. Note that the linked lists must retain their original structure after the function returns.

  -  Custom Judge:
```
     The inputs to the judge are given as follows (your program is not given these inputs):

     intersectVal - The value of the node where the intersection occurs. This is 0 if there is no intersected node.
     listA - The first linked list.
     listB - The second linked list.
     skipA - The number of nodes to skip ahead in listA (starting from the head) to get to the intersected node.
     skipB - The number of nodes to skip ahead in listB (starting from the head) to get to the intersected node.
     The judge will then create the linked structure based on these inputs and pass the two heads, headA and headB to your program. If you correctly return the intersected node, then your solution will be accepted.
```
  
  -  Time Complexity: O(m + n), where m and n are the lengths of the input linked lists headA and headB
   
     Space Complexity: O(1)
     
 - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
    
    Output: Intersected at '8'
    
    Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,6,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B. Note that the intersected node's value is not 1 because the nodes with value 1 in A and B (2nd node in A and 3rd node in B) are different node references. In other words, they point to two different locations in memory, while the nodes with value 8 in A and B (3rd node in A and 4th node in B) point to the same location in memory.
    
    **Example 2:**
    
    Input: intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
    
    Output: Intersected at '2'
    
    Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [1,9,1,2,4]. From the head of B, it reads as [3,2,4]. There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.
    
    
 30. [rotateRight.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/rotateRight.go)

  -  Given the head of a linked list, rotate the list to the right by k places.

  -  The code rotates a linked list to the right by k places by connecting the tail to the head to form a circular list, finding the new tail and new head positions, and breaking the circular list at the new tail position.
  
  -  Time Complexity: O(n), where n is the number of nodes in the linked list
   
     Space Complexity: O(1)
     
 - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: head = [1,2,3,4,5], k = 2
    
    Output: [4,5,1,2,3]

    
    **Example 2:**
    
    Input: head = [0,1,2], k = 4
    
    Output: [2,0,1]

 31. [generateParenthesis.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/generateParenthesis.go)

  -  Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

  -  The code implements a backtracking algorithm to generate all valid combinations of well-formed parentheses. It starts with an empty string and recursively explores all possible choices for adding an opening or closing parenthesis. At each step, it keeps track of the count of open and close parentheses to ensure the generated string remains valid. The base case is reached when the length of the current string reaches 2n (where n is the given input), and at that point, the current combination is added to the result. The algorithm explores all possible combinations by backtracking and returns the final result, which is a list of valid parentheses combinations.
  
  -  Time Complexity: O(4^n / sqrt(n)), where n is the input parameter
   
     Space Complexity: O(n), where n is the input parameter
     
 - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: n = 3
    
    Output: ["((()))","(()())","(())()","()(())","()()()"]

    **Example 2:**
    
    Input: n = 1
    
    Output: ["()"] 
 

32. [divide.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/divide.go)

  -  Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator. The integer division should truncate toward zero, which means losing its fractional part. For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to -2. Return the quotient after dividing dividend by divisor. Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. For this problem, if the quotient is strictly greater than 231 - 1, then return 231 - 1, and if the quotient is strictly less than -231, then return -231.
  
  -  The code implements a division algorithm without using the division operator. It handles special cases and determines the sign of the quotient. The algorithm converts the numbers to positive and uses repeated subtraction to find the quotient. It employs powers of two to efficiently identify the largest multiple of the divisor that can be subtracted from the dividend. Finally, it handles the sign and returns the quotient accordingly.

  -  The code calculates the quotient of dividend divided by divisor using a binary division algorithm, where it repeatedly subtracts the largest multiples of the divisor and accumulates the corresponding powers of two to obtain the quotient.. The algorithm repeatedly finds the largest multiple of the divisor that can be subtracted from the dividend without resulting in a negative value. It does this by doubling the multiple and the corresponding powerOfTwo value until the doubled multiple exceeds the dividend. This identifies the largest power of two that can be subtracted. After finding the largest multiple, it subtracts that multiple from the dividend and adds the corresponding powerOfTwo to the result. This process continues until the dividend becomes smaller than the divisor, indicating that the division is complete.

  
  -  Time Complexity: O(log(dividend/divisor))
   
     Space Complexity: O(1)
     
 - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: dividend = 10, divisor = 3
    
    Output: 3
    
    Explanation: 10/3 = 3.33333.. which is truncated to 3.
    
    **Example 2:**
    
    Input: dividend = 7, divisor = -3
    
    Output: -2
    
    Explanation: 7/-3 = -2.33333.. which is truncated to -2.



33. [searchRangeOn.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/searchRangeOn.go)

  -  Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value. If target is not found in the array, return [-1, -1]. You must write an algorithm with O(n) runtime complexity. This is the O(n) version of the code.
  
  -  The code searches for a given target value in a sorted array of integers. It iterates through the array, storing the indices of matching values in the output slice. The final output is determined based on the number of matching values, returning the starting and ending positions or [-1, -1] if the target is not found.

  - Time Complexity: O(n), where n is the length of the input nums array.
   
     Space Complexity: O(1)
     
 - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: nums = [5,7,7,8,8,10], target = 8
    
    Output: [3,4]
    
    **Example 2:**
   
   Input: nums = [5,7,7,8,8,10], target = 6
   
   Output: [-1,-1]
   
   **Example 3:**
   
   Input: nums = [], target = 0
   
   Output: [-1,-1]

34. [searchRange.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/searchRange.go)

  -  Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value. If target is not found in the array, return [-1, -1]. You must write an algorithm with O(log n) runtime complexity.
  
  -  The provided code implements a binary search algorithm to find the starting and ending positions of a given target value in a sorted array of integers. It utilizes a helper function called binarySearch to perform two separate binary searches: one to find the leftmost occurrence of the target value and another to find the rightmost occurrence. If the target is found, the function returns the range as an array. Otherwise, it returns [-1, -1] to indicate that the target value is not present in the array. The code has a time complexity of O(log n) due to the binary search operations, making it efficient for large arrays.

  - Time Complexity: O(log n), where n is the length of the input nums array.
   
     Space Complexity: O(1)
     
 - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: nums = [5,7,7,8,8,10], target = 8
    
    Output: [3,4]
    
    **Example 2:**
   
   Input: nums = [5,7,7,8,8,10], target = 6
   
   Output: [-1,-1]
   
   **Example 3:**
   
   Input: nums = [], target = 0
   
   Output: [-1,-1]


35. [singleNumber.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/singleNumber.go)

  -  Given a non-empty array of integers nums, every element appears twice except for one. Find that single one. You must implement a solution with a linear runtime complexity and use only constant extra space.

  -   The code is based on the XOR (exclusive OR) operation. XOR returns true (1) if the number of true inputs is odd, and false (0) if the number of true inputs is even. In the given code, we initialize the result variable to 0. By iterating through each element in the nums array, we perform XOR between the current element and the result. This operation has the property that XORing a number with itself results in 0. Therefore, when the same number appears twice, the XOR operation cancels them out, leaving 0 as the result. However, when a number appears only once, the XOR operation with 0 effectively "stores" that number in the result. As we continue iterating through the array, the duplicates cancel out, and the result variable eventually contains the single number that appears only once. Finally, we return the result, which represents the single number that appears only once in the given array.

  -  In the case of 2 XOR 4, the binary representation of these numbers are:
     ```
     2: 0010
     4: 0100   
     0110
     ```
     The resulting binary representation is 0110, which is equal to 6 in decimal.

  - Time Complexity: O(n), where n is the length of the input nums array.
   
     Space Complexity: O(1)
     
 - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: nums = [2,2,1]
    
    Output: 1
    
    **Example 2:**
   
   Input: nums = [4,1,2,1,2]
   
   Output: 4
   
   **Example 3:**
   
   Input: nums = [1]
   
   Output: 1

36. [missingNumber.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/missingNumber.go)

  -  Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

  -   The missingNumber function calculates the expected sum of numbers from 0 to n using the formula totalSum := n * (n + 1) / 2, where n is the length of the input array nums. Then, it iterates through each number in nums and subtracts it from the totalSum. This step effectively cancels out the numbers that are present in nums from the expected sum. Finally, the remaining value of totalSum after the subtraction represents the missing number in the range. This value is returned as the result. By subtracting each number from the expected sum, the function can identify the missing number in the array using the concept of arithmetic progression and the sum of consecutive numbers.

  - Time Complexity: O(n), where n is the length of the input array.
   
     Space Complexity: O(1)
     
 - **Input and Output Examples:**
    
    **Example 1:**
    
    Input: nums = [3,0,1]
    
    Output: 2
    
    Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.

    
    **Example 2:**
   
   Input: nums = [0,1]
   
   Output: 2
   
   Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing number in the range since it does not appear in nums.
   
   **Example 3:**
   
   Input: nums = [9,6,4,2,3,5,7,0,1]
   
   Output: 8
   
   Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums.

36. [findPeakElement.go](https://github.com/bryanoliverh/CodePractice-Golang/blob/main/findPeakElement.go)

  - A peak element is an element that is strictly greater than its neighbors. Given a 0-indexed integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks. You may imagine that nums[-1] = nums[n] = -∞. In other words, an element is always considered to be strictly greater than a neighbor that is outside the array. You must write an algorithm that runs in O(log n) time.

  - At the end of the loop, left and right will have the same value.

  - Time Complexity: O(n), where n is the length of the input array.
   
     Space Complexity: O(1)
     
 - **Input and Output Examples:**
    
    **Example 1:**
    
   Input: nums = [1,2,3,1]
   
   Output: 2
   
   Explanation: 3 is a peak element and your function should return the index number 2.
    
    **Example 2:**
   
   Input: nums = [1,2,1,3,5,6,4]

   Output: 5

   Explanation: Your function can return either index number 1 where the peak element is 2, or index number 5 where the peak element is 6.

   

