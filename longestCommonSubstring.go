func longestCommonSubstring(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    prefix := strs[0]
    for i := 1; i < len(strs); i++ {
        for j := 0; j < len(prefix); j++ {
            for k := len(prefix) - j; k > 0; k-- {
                if strings.Contains(strs[i], prefix[j:j+k]) {
                    prefix = prefix[j : j+k]
                    break
                }
            }
        }
    }
    return prefix
}
