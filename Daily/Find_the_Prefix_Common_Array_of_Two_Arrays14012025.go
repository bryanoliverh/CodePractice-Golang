
func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	C := make([]int, n)
	setA := make(map[int]bool)
	setB := make(map[int]bool)

	for i := 0; i < n; i++ {
		setA[A[i]] = true
		setB[B[i]] = true

		count := 0
		for k := range setA {
			if setB[k] {
				count++
			}
		}
		C[i] = count
	}
	return C
}