func twoSum(array []int, target int) []int { 

seenNums := map[int]int{} 

for i, num := range array { 
  potentialMatch := target - num 
  if j, ok := seenNums[potentialMatch]; ok { 
     return []int{i, j} 
  } 

  seenNums[num] = i 
} 

return []int{} 

}
