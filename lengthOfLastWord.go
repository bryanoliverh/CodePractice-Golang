func lengthOfLastWord(s string) int {
    finArr: = make([] string, 0)
    finStr: = ""
    for i: = 0;i < len(s);i++{
        if s[i] == 32 {
            finStr = ""
        } else {
            finStr += string(s[i])
        }
        if len(finStr) > 0 {
            finArr = append(finArr, finStr)
        }
    }
    if len(finArr) > 0 {
        fmt.Printf("Your last word is \"%s\"\n", finArr[len(finArr) - 1])
        return len(finArr[len(finArr) - 1])
    }
    return 0
}
