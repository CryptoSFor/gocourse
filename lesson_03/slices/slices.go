package slices

func LongestWord(s []string) string {
	max := SymbolsCount(s[0])
	ind := 0
	for i := 1; i < len(s); i++ {
		if SymbolsCount(s[i]) > max {
			max = SymbolsCount(s[i])
			ind = i
		}
	}
	return s[ind]
}

func SymbolsCount(s string) int {
	c := 0
	for range s {
		c += 1
	}
	return c
}

func Reverse(s []int64) []int64 {
	l := len(s)
	r := make([]int64, l)
	for k, v := range s {
		ind := l - k - 1
		r[ind] = v
	}
	return r
}
