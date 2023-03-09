func compress(chars []byte) int { 

if len(chars) == 0 { 

return 0 

} else if len(chars) == 1 { 

return 1 

} 

 
 

comp := chars[0] 

count := 1 

compressIdx := 0 

 
 

for i := 1; i <= len(chars); i++ { 

if i < len(chars) && chars[i] == comp { 

count++ 

} else { 

// chars[i] != comp or i == len(chars) 

chars[compressIdx] = comp 

compressIdx++ 

 
 

if count > 1 { 

countStr := strconv.Itoa(count) 

for j := 0; j < len(countStr); j++ { 

chars[compressIdx+j] = countStr[j] 

} 

compressIdx += len(countStr) 

} 

 
 

if i < len(chars) { 

comp = chars[i] 

} 

count = 1 

} 

} 

 
 

return compressIdx 

} 
