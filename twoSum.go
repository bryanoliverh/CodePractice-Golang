func twoSum(array []int, target int) []int { 

seenNums := map[int]int{} 

for i, num := range array { 

potentialMatch := target - num 

if j, found := seenNums[potentialMatch]; found { 

return []int{i, j} 

} 

seenNums[num] = i 

} 

return []int{} 

} 

 
