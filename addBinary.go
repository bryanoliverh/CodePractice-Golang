func addBinary(a string, b string) string {
    newCarry := 0
    res := ""
    i, j := len(a)-1, len(b)-1
    for i >= 0 || j >= 0 || newCarry > 0 {
        sum := newCarry
        if i >= 0 {
            sum += int(a[i] - '0')
            i--
        }
        if j >= 0 {
            sum += int(b[j] - '0')
            j--
        }
        res = strconv.Itoa(sum%2) + res
        newCarry = sum / 2
    }
    return res
}