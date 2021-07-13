package slices

func MaxLength(slice []string) string {
	max := len(slice[0])
	index := 0
	for i := 1; i < len(slice); i++ {
		if len(slice[i]) > max {
			max = len(slice[i])
			index = i
		}
	}
	return slice[index]
}

func Reverse(slice []int64) []int64 {
	length := len(slice)
	reverse := make([]int64, length)
	for k, v := range slice {
		ind := length - k - 1
		reverse[ind] = v
	}
	return reverse
}
